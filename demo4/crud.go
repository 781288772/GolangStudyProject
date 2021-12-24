package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var sqlDb *sql.DB           //数据库连接db
var sqlResponse SqlResponse //响应client的数据
func init() {
	//1、打开数据库
	//parseTime:时间格式转换(查询结果为时间时，是否自动解析为时间);
	// loc=Local：MySQL的时区设置
	sqlStr := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=true&loc=Local"
	var err error
	sqlDb, err = sql.Open("mysql", sqlStr)
	if err != nil {
		fmt.Println("数据库打开出现了问题：", err)
		return
	}
	//2、 测试与数据库建立的连接（校验连接是否正确）
	err = sqlDb.Ping()
	if err != nil {
		fmt.Println("数据库连接出现了问题：", err)
		return
	}
}

//Client提交的数据
type SqlUser struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

//应答体（响应client的请求）
type SqlResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	r := gin.Default()
	//数据库的CRUD--->gin的 post、get、put、delete方法
	r.POST("sql/insert", insertData)   //添加数据
	r.GET("sql/get", getData)          //查询数据（单条记录）
	r.GET("sql/mulget", getMulData)    //查询数据（多条记录）
	r.PUT("sql/update", updateData)    //更新数据
	r.DELETE("sql/delete", deleteData) //删除数据
	r.Run(":9090")
}

func deleteData(c *gin.Context) {
	name := c.Query("name")
	var count int
	//1、先查询
	sqlStr := "select count(*) from user where name=?"
	err := sqlDb.QueryRow(sqlStr, name).Scan(&count)
	if count <= 0 || err != nil {
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "删除的数据不存在"
		sqlResponse.Data = "error"
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	//2、再删除
	delStr := "delete from user where name=?"
	ret, err := sqlDb.Exec(delStr, name)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "删除失败"
		sqlResponse.Data = "error"
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	sqlResponse.Code = http.StatusOK
	sqlResponse.Message = "删除成功"
	sqlResponse.Data = "OK"
	c.JSON(http.StatusOK, sqlResponse)
	fmt.Println(ret.LastInsertId()) //打印结果
}

func updateData(c *gin.Context) {
	var u SqlUser
	err := c.Bind(&u)
	if err != nil {
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "参数错误"
		sqlResponse.Data = "error"
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	sqlStr := "update user set age=? ,address=? where name=?"
	ret, err := sqlDb.Exec(sqlStr, u.Age, u.Address, u.Name)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "更新失败"
		sqlResponse.Data = "error"
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	sqlResponse.Code = http.StatusOK
	sqlResponse.Message = "更新成功"
	sqlResponse.Data = "OK"
	c.JSON(http.StatusOK, sqlResponse)
	fmt.Println(ret.LastInsertId()) //打印结果
}

func getMulData(c *gin.Context) {
	address := c.Query("address")
	sqlStr := "select name,age from user where address=?"
	rows, err := sqlDb.Query(sqlStr, address)
	if err != nil {
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "查询错误"
		sqlResponse.Data = "error"
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	defer rows.Close()
	resUser := make([]SqlUser, 0)
	for rows.Next() {
		var userTemp SqlUser
		rows.Scan(&userTemp.Name, &userTemp.Age)
		userTemp.Address = address
		resUser = append(resUser, userTemp)
	}
	sqlResponse.Code = http.StatusOK
	sqlResponse.Message = "读取成功"
	sqlResponse.Data = resUser
	c.JSON(http.StatusOK, sqlResponse)
}

func getData(c *gin.Context) {
	name := c.Query("name")
	sqlStr := "select age,address from user where name=?"
	var u SqlUser
	err := sqlDb.QueryRow(sqlStr, name).Scan(&u.Age, &u.Address)
	if err != nil {
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "查询错误"
		sqlResponse.Data = "error"
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	u.Name = name
	sqlResponse.Code = http.StatusOK
	sqlResponse.Message = "读取成功"
	sqlResponse.Data = u
	c.JSON(http.StatusOK, sqlResponse)
}

func insertData(c *gin.Context) {
	var u SqlUser
	err := c.Bind(&u)
	if err != nil {
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "参数错误"
		sqlResponse.Data = "error"
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	sqlStr := "insert into user(name, age, address) values (?,?,?)"
	ret, err := sqlDb.Exec(sqlStr, u.Name, u.Age, u.Address)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		sqlResponse.Code = http.StatusBadRequest
		sqlResponse.Message = "写入失败"
		sqlResponse.Data = "error"
		c.JSON(http.StatusOK, sqlResponse)
		return
	}
	sqlResponse.Code = http.StatusOK
	sqlResponse.Message = "写入成功"
	sqlResponse.Data = "OK"
	c.JSON(http.StatusOK, sqlResponse)
	fmt.Println(ret.LastInsertId()) //打印结果

}

//todo:  go-sql-driver地址：https://github.com/go-sql-driver/mysql
