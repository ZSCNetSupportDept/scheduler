package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(app *echo.Echo) {
	// here is the route for our site
	staticFiles := app.Group("/")
	staticFiles.Use(middleware.Static("./FrontEnd"))
}
