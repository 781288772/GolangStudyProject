package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "root:root@(127.0.0.1:3306)/test")

	defer db.Close()
	err := db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	// fmt.Println("数据库连接成功")
	var id string
	var username string

	rows, err := db.Query("select id,username from sys_user")
	for rows.Next() {

		err := rows.Scan(&id, &username)
		if err != nil {
			log.Fatal("scan failed: ", err)
		}

		log.Printf("id: %s username:%s\n", id, username)
	}

	// fmt.Print(rows)
	if err != nil {
		fmt.Print(err)
		return
	}

	r := gin.Default()

	r.GET("/list", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "查询成功",
			"data":    rows,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
