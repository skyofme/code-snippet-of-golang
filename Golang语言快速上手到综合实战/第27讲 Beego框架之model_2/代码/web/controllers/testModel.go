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
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
	orm.RegisterDataBase("default", "mysql", "root:123@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(UserInfo))

	o := orm.NewOrm()

	//下面是原生读取
/*	var users []UserInfo
	o.Raw("select * from user_info").QueryRows(&users)
	c.Ctx.WriteString(fmt.Sprintf("user info:%v", users))*/


	//采用queryBuilder方式进行读取
	var users []UserInfo

	qb, _:=orm.NewQueryBuilder("mysql")

	qb.Select("password").From("user_info").Where("username= 'lisi'").And("id").In("1").Limit(1)

	sql := qb.String()
	o.Raw(sql).QueryRows(&users)

	c.Ctx.WriteString(fmt.Sprintf("user info:%v", users))
}