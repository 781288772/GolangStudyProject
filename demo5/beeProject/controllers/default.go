package controllers

import (
	"beeProject/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["word"] = "hello world"
	c.TplName = "test.html"
	////获取orm对象
	//o := orm.NewOrm()
	////执行某个操作 增删改查
	//var user models.User
	//user.Name = "小黑"
	//user.Password = "123456"
	////插入操作
	//count, err := o.Insert(&user)
	//if err != nil {
	//	logs.Error("数据插入失败", err)
	//	return
	//}
	//logs.Debug("插入成功")
	//logs.Debug("count", count)
}

func (c *MainController) Post() {
	c.Data["word"] = "Post method"
	c.TplName = "test.html"
}

func (c *MainController) ShowGet() {
	//获取orm对象
	o := orm.NewOrm()
	//执行某个操作 增删改查
	var user models.User
	user.Name = "小黑"
	user.Password = "123456"
	//插入操作
	count, err := o.Insert(&user)
	if err != nil {
		logs.Error("数据插入失败", err)
		return
	}
	logs.Debug("插入成功")
	logs.Debug("count", count)

}
func (c *MainController) UpdateUser() {
	o := orm.NewOrm()
	var user models.User
	user.Id = 1
	user.Name = "XIUGA"

	count, err := o.Update(&user)
	if err != nil {
		logs.Error("查询数据失败")
		return
	}
	logs.Debug("count", count)

}
