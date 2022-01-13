package models

import "github.com/astaxie/beego/orm"
import (
	_ "github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

//表的设计

//定义一个结构体
type User struct {
	Id       int
	Name     string
	Password string
}

func init() {
	//获取连接对象
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	//创建表
	orm.RegisterModel(new(User))
	//生成表
	//第一个参数，数据库别名，第二个参数是否强制更新表
	orm.RunSyncdb("default", false, true)

}
