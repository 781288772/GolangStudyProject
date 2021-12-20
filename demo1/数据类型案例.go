package main

import "fmt"

func main() {
	//	请输入用户姓名，年龄，当输入完成以后在屏幕显示：您好，XX你的年龄为XX

	//1.定义变量
	var name string
	var age int

	//2.给出相应的录入提示
	fmt.Printf("请输入姓名:")
	fmt.Scan(&name)
	fmt.Println("请输入年龄:")
	fmt.Scan(&age)

	fmt.Printf("您好%s,您的年龄为%d", name, age)
}
