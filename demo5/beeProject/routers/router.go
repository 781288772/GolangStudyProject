package routers

import (
	"beeProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	////给请求指定自定义方法，一个请求一个方法
	//beego.Router("/login", &controllers.LoginController{}, "get:showLogin;post:postFunc")
	////多个请求一个方法
	//beego.Router("/index", &controllers.LoginController, "get,post:handleFunc")
	////给所有请求指定方法
	//beego.Router("/index", &controllers.LoginController, "*:handleFunc")

}
