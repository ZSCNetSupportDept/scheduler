package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/labstack/echo/v4"
	"zsxyww.com/scheduler/config"
	db "zsxyww.com/scheduler/database"
	"zsxyww.com/scheduler/model"
	"zsxyww.com/scheduler/route"
	tl "zsxyww.com/scheduler/templates"
)

func main() {

	//进行各种初始化工作：

	config.Load()
	db.Connect()

	app := echo.New()
	csv() //初始化Model.MemberList

	route.Route(app)      //注册路由表
	route.Middleware(app) //注册中间件

	// 暂时在初始化时不注册模板，因为用不上
	//registerTemplate(app) //注册模板

	listenAddress := fmt.Sprintf(":%d", config.Default.App.ListenPort)

	app.Logger.Fatal(app.Start(listenAddress)) //启动服务器
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

func registerTemplate(app *echo.Echo) {
	renderer := tl.Tlw{
		Tl: template.Must(template.ParseGlob("templates/*.html")),
	}
	app.Renderer = renderer
}
