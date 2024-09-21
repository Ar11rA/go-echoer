package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	// load environment
	loadEnv()

	// new echo server
	e := echo.New()

	// register middleware and services
	registerMiddleware(e)
	registerServices(e)
	e.Logger.Fatal(e.Start(":7001"))
}
