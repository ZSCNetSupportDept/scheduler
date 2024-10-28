package main

import (
	"github.com/labstack/echo/v4"
	"zsxyww.com/scheduler/config"
	"zsxyww.com/scheduler/database"
)

func main() {

	config.Load()
	db.Connect()

	app := echo.New()
	register(app)
}
func register(app *echo.Echo) {
	config.Route(app)
	config.Middleware(app)
}
