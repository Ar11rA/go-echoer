package routes

import (
	"net/http"
	"quote-server/services"
	"quote-server/types"

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
func httpGetHandler(c echo.Context, httpService services.HttpService) error {
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

// postHandler handles the POST request to echo the body back
// @Summary POST Handler
// @Description Echoes the posted body back from Postman Echo
// @Accept  json
// @Produce json
// @Param data body types.EchoRequest true "Post Request Body"
// @Success 200 {object} types.EchoResponse "Successful response"
// @Failure 400 {string} string "Bad request - invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /echo [post]
func httpPostHandler(c echo.Context, httpService services.HttpService) error {
	var echoRequest types.EchoRequest
	if err := c.Bind(&echoRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	// Call the service
	response, err := httpService.PostEcho(echoRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}
