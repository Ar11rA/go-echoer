package routes

import (
	"quote-server/services"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes registers all routes with the Echo instance
func RegisterRoutes(e *echo.Echo, container services.Container) {
	e.GET("/health", healthHandler)
	e.GET("/env", envHandler)
	e.GET("/save", func(c echo.Context) error {
		return fileHandler(c, container.FileService)
	})
	e.GET("/echo", func(c echo.Context) error {
		return httpHandler(c, container.HttpService)
	})
}
