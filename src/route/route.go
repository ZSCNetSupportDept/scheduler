package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"zsxyww.com/scheduler/config"
	"zsxyww.com/scheduler/handler"
)

func Route(app *echo.Echo) {
	// here is the route for our site
	staticFiles := app.Group("/*")
	staticFiles.Use(middleware.Static(config.Default.App.FrontEndDir))

	api := app.Group("/api/")
	api.GET("getAssignment", handler.GetAssignment)
	api.GET("test", handler.HandlerTest)
}
