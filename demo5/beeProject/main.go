package main

import (
	_ "beeProject/models"
	_ "beeProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
