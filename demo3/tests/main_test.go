package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"go-api/config"
)

func setup() {
	gin.SetMode(gin.TestMode)
	//设置config模式
	config.SetConfigMode("test")
	config.InitConfig()
	//打印config
	fmt.Println(config.AppSetting.JwtSecret)
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}
func TestMain(m *testing.M) {
	setup()
	fmt.Println("Test begins....")
	code := m.Run() // 如果不加这句，只会执行Main
	teardown()
	os.Exit(code)
}
