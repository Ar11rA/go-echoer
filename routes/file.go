package routes

import (
	"fmt"
	"net/http"
	"quote-server/services"

	"github.com/labstack/echo/v4"
)

func fileHandler(c echo.Context, fileService services.FileService) error {
	content := c.QueryParam("content")
	if content == "" {
		return c.String(http.StatusBadRequest, "Content is required")
	}
	if err := fileService.Save(content); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Failed to save content")
	}
	return c.String(http.StatusOK, "Content saved successfully")
}
