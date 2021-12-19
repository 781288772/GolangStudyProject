package main

import "fmt"

func main() {
	var age int
	fmt.Printf("请输入年龄")
	//fmt.Scanf("%d", &age) //通过scanf函数将键盘输入的数据赋值给变量，但变量前必须加取地址符号&
	//fmt.Println("age=", age)

	fmt.Scan(&age)
	fmt.Print("age=", age)

}
