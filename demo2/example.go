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
type sqlResponse struct {
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
		log.Print("data", resUser)

		// fmt.Print(rows)
		if err != nil {
			fmt.Print(err)
			return
		}
		c.JSON(http.StatusOK, resUser)

		// c.JSON(http.StatusOK, gin.H{
		// 	"status":  "200",
		// 	"message": "查询成功",
		// 	"data":    resUser,
		// })
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
