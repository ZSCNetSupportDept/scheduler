package main

import (
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
	if config.Default.Option.DatabaseAutoMigrate == true {
		db.Main.AutoMigrate(&model.Member{}, &model.Tweak{})
	}

	app := echo.New()
	csv() //初始化Model.MemberList

	route.Route(app)      //注册路由表
	route.Middleware(app) //注册中间件

	renderer := tl.Tlw{
		Tl: template.Must(template.ParseGlob(config.Default.App.TemplateDir + "/*.html")),
	}
	app.Renderer = renderer //注册模板

	app.Logger.Fatal(app.Start(config.Default.App.ListenPath)) //启动服务器
}

// 读取csv文件
func csv() {
	data, err := os.OpenFile(config.Default.App.MemberFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer data.Close()

	err = gocsv.UnmarshalFile(data, &model.MemberList)
	if err != nil {
		panic(err)
	}
}
