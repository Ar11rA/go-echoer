package routes

import (
	"fmt"
	"net/http"
	"quote-server/services"

	"github.com/labstack/echo/v4"
)

// @Summary File Handler
// @Description Saves content to a file based on the provided query parameter
// @Produce text/plain
// @Param content query string true "Content to save"
// @Success 200 {string} string "Content saved successfully"
// @Failure 400 {string} string "Bad request - content is required"
// @Failure 500 {string} string "Internal server error - failed to save content"
// @Router /save [get]
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
