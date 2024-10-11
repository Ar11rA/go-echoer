package main

import (
	"context"
	"fmt"
	"os"
	"quote-server/config"
	"quote-server/routes"
	"quote-server/services"
	"quote-server/utils"
	"strconv"
	"time"

	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Info("Environments not loaded using .env file")
	}
}

func registerMiddleware(e *echo.Echo) {
	e.Logger.SetOutput(config.LogrusLogger.Out)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: config.LogrusLogger.Out,
	}))
	e.Use(middleware.Recover()) // Recovers from panic
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Allows all origins
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
}

func registerServices(e *echo.Echo) {

	// Initialize other services here as needed
	container := services.Container{
		FileService:  registerFileService(),
		HttpService:  registerHttpService(),
		DBService:    registerDBService(),
		RedisService: registerCacheService(),
		MQService:    registerMQService(),
		LogService:   registerMongoService(),
		KafkaService: registerKafkaService(),
	}

	// Register routes
	routes.RegisterRoutes(e, container)
}

func registerFileService() services.FileService {
	dir := os.Getenv("WRITE_PATH")
	if dir == "" {
		dir = "./" // default directory
	}

	// Create a new instance of OSFileWriter
	fileWriter := &utils.OSFileWriter{}

	// Create a new instance of FileServiceImpl
	fileService := services.NewFileService(dir, fileWriter)

	return fileService
}

func registerHttpService() services.HttpService {
	httpClient := &utils.EchoHttpClient{}
	httpService := services.NewHttpService(httpClient)

	return httpService
}

func registerDBService() services.DBService {
	portStr := os.Getenv("PG_DB_PORT") // Read port from environment variable
	port, err := strconv.Atoi(portStr) // Convert string to int
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}

	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     os.Getenv("PG_DB_HOST"),
		Port:     uint16(port),
		User:     os.Getenv("PG_DB_USER"),
		Password: os.Getenv("PG_DB_PASS"),
		Database: os.Getenv("PG_DB_NAME"),
	})
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v", err))
	}

	dbClient := &utils.PgxClient{Conn: conn}
	dbService := services.NewDbService(dbClient)
	return dbService
}

func registerCacheService() services.RedisService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})

	redisClient := &utils.RedisUtil{Client: rdb}
	redisService := services.NewRedisService(redisClient)
	return redisService
}

func registerMQService() services.RabbitMQService {
	conn, err := amqp091.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	// Create RabbitMQ utility
	rabbitUtil, err := utils.NewRabbitMQUtil(conn)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ util: %v", err)
	}

	// Create RabbitMQ service
	rabbitMQService := services.NewRabbitMQService(rabbitUtil)
	return rabbitMQService
}

func registerMongoService() services.LogService {
	// Set MongoDB connection options
	clientOptions := options.Client().
		ApplyURI(os.Getenv("MONGODB_URI")) // You should set this environment variable

	// Create a new client and connect to the MongoDB server
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the database to verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	// Create the MongoDB utility with specified database and collection names
	dbName := os.Getenv("MONGODB_DB_NAME") // Set this environment variable
	mongoUtil := utils.NewMongoDBUtil(client, dbName)

	// Create the LogService using the MongoDB utility
	logService := services.NewLogService(mongoUtil)
	return logService
}

func registerKafkaService() services.KafkaService {
	// Get the Kafka broker address from the environment variable
	brokerAddress := os.Getenv("KAFKA_BROKER_ADDRESS") // Set this environment variable
	topic := os.Getenv("KAFKA_TOPIC")                  // Set this environment variable

	// Create a Kafka writer with necessary configurations
	kafkaWriter := &kafka.Writer{
		Addr:     kafka.TCP(brokerAddress),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	// Optionally, you can create a Kafka reader or consumer here for bi-directional communication
	// e.g., kafkaReader := kafka.NewReader(kafka.ReaderConfig{...})

	// Create the Kafka utility
	kafkaUtil := utils.NewKafkaUtil(kafkaWriter)

	// Create the KafkaService using the Kafka utility
	kafkaService := services.NewKafkaService(kafkaUtil)
	return kafkaService
}
