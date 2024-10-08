package routes

import (
	"net/http"
	"quote-server/config"
	"quote-server/services"
	"quote-server/types"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary Insert a new log
// @Description Inserts a new log entry into the database
// @Accept json
// @Produce json
// @Param log body types.LogRequest true "Log entry"
// @Success 201 {object} types.Log "Log successfully created"
// @Failure 400 {string} string "Bad request - invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /logs [post]
func postLogHandler(c echo.Context, logService services.LogService) error {
	var logEntry types.LogRequest
	if err := c.Bind(&logEntry); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	// Call the service to insert the log
	config.LogrusLogger.Info("Application id: " + logEntry.ApplicationID)
	if err := logService.InsertLog(logEntry.ApplicationID, logEntry.Logs); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, logEntry)
}

// @Summary Get logs by application ID
// @Description Fetches logs from the database by application ID
// @Produce json
// @Param appID query string true "Application ID"
// @Param limit query int false "Number of logs to fetch"
// @Success 200 {array} types.Log "List of logs"
// @Failure 400 {string} string "Bad request - application ID is required"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /logs [get]
func getLogsHandler(c echo.Context, logService services.LogService) error {
	appID := c.QueryParam("appID")
	if appID == "" {
		return c.JSON(http.StatusBadRequest, "application ID is required")
	}

	// Get the limit from the query parameter (optional)
	limitParam := c.QueryParam("limit")
	limit := 10 // Default limit
	if limitParam != "" {
		var err error
		limit, err = strconv.Atoi(limitParam)
		if err != nil || limit <= 0 {
			return c.JSON(http.StatusBadRequest, "invalid limit parameter")
		}
	}

	// Call the service to fetch logs
	logs, err := logService.GetLogsByApplicationID(appID, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, logs)
}
