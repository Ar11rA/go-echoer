package routes

import (
	"context"
	"fmt"
	"net/http"
	"quote-server/config"
	"quote-server/services"
	"quote-server/types"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// @Summary Publish Message to RabbitMQ
// @Description Publishes a message to a RabbitMQ queue
// @Accept application/json
// @Param message body types.MessagePublishRequest true "MessagePublishRequest object"
// @Produce application/json
// @Success 200 {string} string "Message published successfully"
// @Failure 400 {string} string "Bad request - Content is required"
// @Failure 500 {string} string "Failed to publish message"
// @Router /rabbit/publish [post]
func rabbitMQPostHandler(c echo.Context, rabbitMQService services.RabbitMQService) error {
	var req types.MessagePublishRequest

	// Parse request body
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}

	// Validate request
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Convert content to byte array
	message := []byte(req.Content)
	ctx := context.Background()

	// Publish message to RabbitMQ
	if err := rabbitMQService.SendMessage(ctx, message); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send message"})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "Message sent successfully"})
}

// @Summary Publish Message to Kafka
// @Description Publishes a message to a Kafka topic using key and message from path params
// @Accept json
// @Produce json
// @Param key path string true "Key for Kafka message"
// @Param message path string true "Message to be published"
// @Success 200 {string} string "Message published successfully"
// @Failure 400 {string} string "Bad request - Missing key or message"
// @Failure 500 {string} string "Failed to publish message"
// @Router /kafka/publish/{key}/{message} [get]
func kafkaGetHandler(c echo.Context, kafkaService services.KafkaService) error {
	// Get key and message from path params
	key := c.Param("key")
	message := c.Param("message")

	// Validate that key and message are provided
	if key == "" || message == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Key and message are required"})
	}

	// Publish the message to Kafka
	ctx := context.Background()
	if err := kafkaService.PublishMessage(ctx, key, message); err != nil {
		config.LogrusLogger.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send message"})
	}

	config.LogrusLogger.WithFields(logrus.Fields{
		"key":     key,
		"message": message,
	}).Info("Message published")

	// Return success response
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
