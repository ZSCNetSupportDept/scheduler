package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/labstack/echo/v4"
	"html/template"
	"os"
	"zsxyww.com/scheduler/config"
	"zsxyww.com/scheduler/database"
	"zsxyww.com/scheduler/model"
	"zsxyww.com/scheduler/route"
	"zsxyww.com/scheduler/templates"
)

func main() {

	config.Load()
	db.Connect()

	app := echo.New()
	register(app)
	csv()

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

// 读取csv文件
func csv() {
	data, err := os.OpenFile(config.Default.App.File, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer data.Close()

	err = gocsv.UnmarshalFile(data, &model.MemberList)
	if err != nil {
		panic(err)
	}
}
