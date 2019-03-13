package routers

import (
	"web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/test_input", &controllers.TestInputController{}, "get:Get;post:Post")
}
