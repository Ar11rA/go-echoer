// routes/redis_routes.go
package routes

import (
	"context"
	"net/http"
	"quote-server/services"
	"quote-server/types"

	"github.com/labstack/echo/v4"
)

// @Summary Save Data to Redis
// @Description Save a key-value pair to Redis using POST
// @Accept  json
// @Produce text/plain
// @Param data body types.RedisDataRequest true "Key-value pair to save"
// @Success 200 {string} string "Data saved successfully"
// @Failure 400 {string} string "Bad request - invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /redis/save [post]
func redisSaveHandler(c echo.Context, redisService services.RedisService) error {
	var req types.RedisDataRequest

	// Bind and validate the request body
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	if err := c.Validate(&req); err != nil {
		return c.String(http.StatusBadRequest, "Key and value are required")
	}

	// Create a background context
	ctx := context.Background()

	// Save data using the Redis service
	err := redisService.SaveData(ctx, req.Key, req.Value)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save data")
	}

	return c.String(http.StatusOK, "Data saved successfully")
}

// @Summary Get Data from Redis
// @Description Retrieve a value from Redis using the provided key
// @Produce json
// @Param key path string true "Key to retrieve value"
// @Success 200 {object} map[string]string "Value retrieved successfully"
// @Failure 500 {string} string "Internal server error - failed to retrieve data"
// @Router /redis/{key} [get]
func redisGetHandler(c echo.Context, redisService services.RedisService) error {
	key := c.Param("key")
	ctx := context.Background()

	value, err := redisService.GetData(ctx, key)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve data")
	}

	return c.JSON(http.StatusOK, map[string]string{"value": value})
}
