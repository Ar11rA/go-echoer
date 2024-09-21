package routes

import (
	"net/http"
	"os"
	"quote-server/config"

	"github.com/labstack/echo/v4"
)

// @Summary Health Check
// @Description Check the health of the server
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func healthHandler(c echo.Context) error {
	config.LogrusLogger.Info("Within health route")
	return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
}

// @Summary Sample env value Check
// @Description Check the env vars of the server
// @Produce json
// @Success 200 {object} map[string]string
// @Router /env [get]
func envHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"limit": os.Getenv("QUOTE_LIMIT")})
}
