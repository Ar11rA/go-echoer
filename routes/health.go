package routes

import (
	"net/http"
	"os"
	"quote-server/config"

	"github.com/labstack/echo/v4"
)

func healthHandler(c echo.Context) error {
	config.LogrusLogger.Info("Within health route")
	return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
}

func envHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"limit": os.Getenv("QUOTE_LIMIT")})
}
