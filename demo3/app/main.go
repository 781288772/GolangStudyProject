package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(config.RunMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// r.Run(fmt.Sprintf("%s:%d", config.ServerSetting.HttpAddress, config.ServerSetting.HttpPort)) // listen and serve on 0.0.0.0:8080
	r.Run()
}
