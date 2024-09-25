package routes

import (
	"net/http"
	"quote-server/services"
	"quote-server/types"
	"strconv"

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

// @Summary HTTP Handler for Fetching Multiple Quotes
// @Description Fetches multiple quotes from Postman Echo API based on the limit parameter
// @Produce json
// @Param limit query int true "Number of quotes to fetch"
// @Success 200 {array} types.QuoteResponse "List of quotes"
// @Failure 400 {string} string "Bad request - limit parameter is required or invalid"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /quotes [get]
func httpGetQuotesHandler(c echo.Context, httpService services.HttpService) error {
	// Get the limit from the query parameter
	limitParam := c.QueryParam("limit")
	if limitParam == "" {
		return c.JSON(http.StatusBadRequest, "limit parameter is required")
	}

	// Convert the limit to an integer
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		return c.JSON(http.StatusBadRequest, "invalid limit parameter")
	}

	// Call the service to fetch quotes
	quotes, err := httpService.GetQuotes(int32(limit))

	// If there were any partial errors, we include them in the response
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Return the successful list of quotes
	return c.JSON(http.StatusOK, quotes)
}
