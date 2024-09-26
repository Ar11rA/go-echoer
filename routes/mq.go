package routes

import (
	"context"
	"fmt"
	"net/http"
	"quote-server/services"
	"quote-server/types"

	"github.com/labstack/echo/v4"
)

// @Summary Publish Message to RabbitMQ
// @Description Publishes a message to a RabbitMQ queue
// @Accept application/json
// @Param message body types.MessagePublishRequest true "MessagePublishRequest object"
// @Produce application/json
// @Success 200 {string} string "Message published successfully"
// @Failure 400 {string} string "Bad request - Content is required"
// @Failure 500 {string} string "Failed to publish message"
// @Router /publish [post]
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
