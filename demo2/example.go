package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	initDB()
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello gin",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func initDB() {
	db, _ := sql.Open("mysql", "root:root@(127.0.0.1:3306)/test")

	defer db.Close()
	err := db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	fmt.Println("数据库连接成功")
	row, err := db.Query("select *from sys_user")
	fmt.Print(row)

}
