package main

import (
	"fmt"
	"os"
	"quote-server/config"
	"quote-server/routes"
	"quote-server/services"
	"quote-server/utils"
	"strconv"

	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
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
		FileService: registerFileService(),
		HttpService: registerHttpService(),
		DBService:   registerDBService(),
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
