package main

import (
	"quote-server/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	// load environment
	loadEnv()

	// new echo server
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	// register middleware and services
	registerMiddleware(e)
	registerServices(e)
	e.Logger.Fatal(e.Start(":7001"))
}
