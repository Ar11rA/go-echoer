package routes

import (
	"net/http"
	"quote-server/services"

	"github.com/labstack/echo/v4"
)

// @Summary HTTP Handler
// @Description Fetches data from Postman Echo based on a query parameter
// @Produce json
// @Param query query string true "Query parameter"
// @Success 200 {object} types.EchoResponse "Successful response"
// @Failure 400 {string} string "Bad request - query parameter is required"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /echo [get]
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
