package routes

import (
	"net/http"
	"quote-server/services"

	"github.com/labstack/echo/v4"
)

func httpHandler(c echo.Context, httpService services.HttpService) error {
	query := c.QueryParam("query")
	if query == "" {
		return c.JSON(http.StatusBadRequest, "query parameter is required")
	}

	// Call the service
	response, err := httpService.GetEcho(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}
