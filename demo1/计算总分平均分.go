package main

import "fmt"

func main() {

	var chinese int
	var math int
	var english int

	fmt.Println("请输入语文成绩")
	fmt.Scan(&chinese)
	fmt.Println("请输入数学成绩")
	fmt.Scan(&math)
	fmt.Println("请输入英语成绩")
	fmt.Scan(&english)

	sum := chinese + english + math
	ava := sum / 3
	fmt.Printf("你的成绩总分是%d,平均分是%d", sum, ava)

}
