package models

import _ "github.com/go-sql-driver/mysql"

func init() {
	//操作数据库
	//sqlStr := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true&loc=Local"
	//conn, err := sql.Open("mysql", sqlStr)
	//if err != nil {
	//	logs.Error("连接错误", err)
	//	return
	//}
	////关闭数据库
	//defer conn.Close()
	////_, err = conn.Exec("create table user(name varchar(40) ,password varchar(40));")
	////if err != nil {
	////	logs.Error("创建表失败", err)
	////}
	////conn.Exec("insert user(name,password) value (?,?)","小黑","123456")
	//
	//res, _ := conn.Query("select name from user")
	//var name string
	//for res.Next() {
	//	res.Scan(&name)
	//	logs.Info("name", name)
	//}

}
