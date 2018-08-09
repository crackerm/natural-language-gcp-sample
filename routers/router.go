package routers

import (
	"github.com/astaxie/beego"
	"testProject/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
