package routers

import (
	"web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/test_httplib", &controllers.TestHttpLibController{}, "get:Get;post:Post")
    beego.Router("/test_context", &controllers.TestContextController{}, "get:Get;post:Post")
}
