package routes

import (
	"quote-server/types"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes registers all routes with the Echo instance
func RegisterRoutes(e *echo.Echo, container types.Container) {
	e.GET("/health", healthHandler)
	e.GET("/env", envHandler)
	e.GET("/save", func(c echo.Context) error {
		return fileHandler(c, container.FileService)
	})
}
