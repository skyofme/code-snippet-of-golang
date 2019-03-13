package main

import (
	"github.com/astaxie/beego"
)

/* 忘记是干嘛的了。。。 */
type HomeController struct {
	beego.Controller
}

// 重写 Get方法
func (this *HomeController) Get() {
	this.Ctx.WriteString("hello")
}

func main() {
	// 注册路由
	beego.Router("/", &HomeController{})
	// 启动服务器
	beego.Run()
}
