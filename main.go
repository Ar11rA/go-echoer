package main

import (
	"os"
	"quote-server/config"
	"quote-server/routes"
	"quote-server/services"
	"quote-server/types"
	"quote-server/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Info("Environments not loaded using .env file")
	}

	e := echo.New()
	e.Logger.SetOutput(config.LogrusLogger.Out)

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: config.LogrusLogger.Writer(),
	}))
	e.Use(middleware.Recover()) // Recovers from panic

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Allows all origins
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	dir := os.Getenv("WRITE_PATH")
	if dir == "" {
		dir = "./" // default directory
	}

	// Create a new instance of OSFileWriter
	fileWriter := &utils.OSFileWriter{}

	// Create a new instance of FileServiceImpl
	fileService := &services.FileServiceImpl{
		Directory:  dir,
		FileWriter: fileWriter,
	}

	container := types.Container{
		FileService: fileService,
		// Initialize other services here as needed
	}

	// Register routes
	routes.RegisterRoutes(e, container)

	e.Logger.Fatal(e.Start(":7001"))
}
