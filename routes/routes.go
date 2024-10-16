package routes

import (
	"quote-server/services"

	_ "quote-server/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// RegisterRoutes registers all routes with the Echo instance
func RegisterRoutes(e *echo.Echo, container services.Container) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/health", healthHandler)
	e.GET("/env", envHandler)
	e.GET("/save", func(c echo.Context) error {
		return fileHandler(c, container.FileService)
	})
	e.GET("/echo", func(c echo.Context) error {
		return httpGetHandler(c, container.HttpService)
	})
	e.GET("/quotes", func(c echo.Context) error {
		return httpGetQuotesHandler(c, container.HttpService)
	})
	e.POST("/echo", func(c echo.Context) error {
		return httpPostHandler(c, container.HttpService)
	})
	e.POST("/users", func(c echo.Context) error {
		return createUserHandler(c, container.DBService)
	})
	e.GET("/users/:id", func(c echo.Context) error {
		return getUserByIDHandler(c, container.DBService)
	})
	e.POST("/redis/save", func(c echo.Context) error {
		return redisSaveHandler(c, container.RedisService)
	})
	e.GET("/redis/:key", func(c echo.Context) error {
		return redisGetHandler(c, container.RedisService)
	})
	e.POST("/rabbit/publish", func(c echo.Context) error {
		return rabbitMQPostHandler(c, container.MQService)
	})
	e.GET("/kafka/publish/:key/:message", func(c echo.Context) error {
		return kafkaGetHandler(c, container.KafkaService)
	})
	e.POST("/logs", func(c echo.Context) error {
		return postLogHandler(c, container.LogService)
	})
	e.GET("/logs", func(c echo.Context) error {
		return getLogsHandler(c, container.LogService)
	})
}
