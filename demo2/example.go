package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}
type SqlResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	db, _ := sql.Open("mysql", "root:root@(127.0.0.1:3306)/test")

	defer db.Close()
	err := db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	// fmt.Println("数据库连接成功")
	// var uid string
	// var uusername string

	r := gin.Default()

	// 查询所有用户
	r.GET("/list", func(c *gin.Context) {
		resUser := make([]User, 0)
		rows, err := db.Query("select id,username from sys_user")

		for rows.Next() {
			var userTemp User
			err := rows.Scan(&userTemp.Id, &userTemp.Username)
			if err != nil {
				log.Fatal("scan failed: ", err)
			}
			log.Printf("id: %s username:%s\n", userTemp.Id, userTemp.Username)

			// log.Print("temp", userTemp)
			resUser = append(resUser, userTemp)
		}
		// log.Print("data", resUser)

		fmt.Print(resUser)
		if err != nil {
			fmt.Print(err)
			return
		}
		var sqlResponse SqlResponse
		sqlResponse.Code = http.StatusOK
		sqlResponse.Message = "操作成功"
		sqlResponse.Data = resUser

		c.JSON(http.StatusOK, sqlResponse)

		// c.JSON(http.StatusOK, gin.H{
		// 	"status":  "200",
		// 	"message": "查询成功",
		// 	"data":    resUser,
		// })
	})
	// 添加用户
	// 路径传参形式
	r.POST("/register", func(c *gin.Context) {

		username := c.PostForm("username")
		password := c.PostForm("password")
		log.Print("username", username)
		log.Print("password", password)
		res, err := db.Exec("INSERT INTO sys_user (username, password) VALUES (?,?)", username, password)
		if err != nil {
			log.Print("插入数据失败")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "操作成功",
			"data":    res,
		})
		// db.Exec("")

	})
	// 更新用户信息
	r.POST("/update", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		email := c.PostForm("email")
		id := c.PostForm("id")
		log.Print("username", username)
		log.Print("password", password)
		log.Print("email", email)
		log.Print("id", id)
		res, err := db.Exec("UPDATE  sys_user SET username = ?, password=?,email=? where id = ?", username, password, email, id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "500",
				"message": "操作失败",
				"data":    res,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "操作成功",
			"data":    res,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
