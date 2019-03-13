package routers

import (
	"web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/test_view", &controllers.TestViewController{}, "get:Get;post:Post")
}
