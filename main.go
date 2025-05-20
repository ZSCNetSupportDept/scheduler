package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"zsxyww.com/scheduler/config"
	"zsxyww.com/scheduler/database"
	"zsxyww.com/scheduler/route"
	"zsxyww.com/scheduler/templates"
)

func main() {

	config.Load()
	db.Connect()

	app := echo.New()
	register(app)

	listenAddress := fmt.Sprintf(":%d", config.Default.App.ListenPort)

	app.Logger.Fatal(app.Start(listenAddress))
}
func register(app *echo.Echo) {
	route.Route(app)
	route.Middleware(app)
	renderer := tl.Tlw{
		Tl: template.Must(template.ParseGlob("templates/*.html")),
	}
	app.Renderer = renderer

}
