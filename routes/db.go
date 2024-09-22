package routes

import (
	"fmt"
	"net/http"
	"quote-server/services"
	"quote-server/types"
	"strconv"

	"github.com/jackc/pgx"
	"github.com/labstack/echo/v4"
)

// @Summary Create User
// @Description Creates a new user with the provided username and email
// @Produce json
// @Param user body types.UserRequest true "User object"
// @Success 201 {object} map[string]bool "User created successfully"
// @Failure 400 {string} string "Bad request - username and email are required"
// @Failure 500 {string} string "Internal server error - failed to create user"
// @Router /users [post]
func createUserHandler(c echo.Context, dbService services.DBService) error {
	var user types.UserRequest

	// Bind the incoming JSON to the user struct
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	// Create the user
	err := dbService.InsertUser(types.UserModel{
		Username: user.Name,
		Email:    user.Email,
	})

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusCreated, map[string]bool{"status": true})
}

// @Summary Get User by ID
// @Description Retrieves user information by user ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} types.UserModel "User found"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error - failed to retrieve user"
// @Router /users/{id} [get]
func getUserByIDHandler(c echo.Context, dbService services.DBService) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 32) // Convert ID to int32
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}
	// Get the user by ID
	user, err := dbService.GetUserById(int32(id))
	if err != nil {
		fmt.Println(err)
		if err == pgx.ErrNoRows {
			return c.String(http.StatusNotFound, "User not found")
		}
		return c.String(http.StatusInternalServerError, "Failed to retrieve user")
	}

	return c.JSON(http.StatusOK, user)
}
