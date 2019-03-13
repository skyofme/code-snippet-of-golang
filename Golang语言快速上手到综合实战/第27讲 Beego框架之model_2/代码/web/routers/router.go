package routers

import (
	"web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/test_model", &controllers.TestModelController{}, "get:Get;post:Post")
}
