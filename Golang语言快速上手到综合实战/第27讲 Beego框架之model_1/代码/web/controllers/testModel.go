package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

//由于model这个名字叫 UserInfo 那么操作的表其实 user_info
type UserInfo struct{ 
	Id int64
	Username string
	Password string
}

type TestModelController struct {
	beego.Controller
}

func (c *TestModelController) Get() {
	orm.RegisterDataBase("default", "mysql", "root:123@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(UserInfo))

	o := orm.NewOrm()
	//下面是插入
	//user := UserInfo{Username:"zhangsan", Password:"123456"}
	//id,err := o.Insert(&user)
	//下面是更新
	//user := UserInfo{Username:"zhangsan", Password:"123456"}
	//user.Id = 1
	//o.Update(&user)
	
	//下面是读取
	user := UserInfo{Id:1}
	o.Read(&user)

	c.Ctx.WriteString(fmt.Sprintf("user info:%v", users))
}