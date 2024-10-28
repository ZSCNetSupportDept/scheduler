package config

import (
	"github.com/labstack/echo/v4"
)

func Route(app *echo.Echo) {
	// here is the route for our site
	app.File("/", "FrontEnd/index.html")
}
