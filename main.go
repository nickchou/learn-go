package main

import (
	"github.com/nickchou/learn-go/app"
	"github.com/nickchou/learn-go/controller"
)

func main() {

	app.Static["/js"] = "static/js"
	app.AutoRouter(&controller.IndexController{})
	app.AutoRouter(&controller.CommController{})
	app.AutoRouter(&controller.CalendarController{})
	app.AutoRouter(&controller.ZhuzherController{})
	app.AutoRouter(&controller.StringController{})
	app.AutoRouter(&controller.AreaController{})
	app.Router("login", &controller.LoginController{})
	app.RunOn(":9090")
}
