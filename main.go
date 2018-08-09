package main

import (
	"github.com/astaxie/beego"
	"testProject/controllers"
	_ "testProject/routers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/testnatural/", &controllers.TestNaturalLanguageController{}, "get:GetResult;post:ProcessText")
	beego.Run()
}
