package services

type Container struct {
	FileService  FileService
	HttpService  HttpService
	DBService    DBService
	RedisService RedisService
	MQService    RabbitMQService
}
