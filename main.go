package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"zsxyww.com/scheduler/config"
	"zsxyww.com/scheduler/database"
	"zsxyww.com/scheduler/route"
)

func main() {

	config.Load()
	db.Connect()

	app := echo.New()
	register(app)

	listenAddress := fmt.Sprintf(":%d", config.ListenPort)

	app.Logger.Fatal(app.Start(listenAddress))
}
func register(app *echo.Echo) {
	route.Route(app)
	route.Middleware(app)
}
