package main

import (
	"fmt"
)

func main() {
	//	输入半径
	var r float64
	const PI = 3.145
	fmt.Println("请输入圆的半径")
	fmt.Scan(&r)

	//计算圆的面积
	s := PI * r * r

	//计算圆的周长

	c := 2 * r * PI

	fmt.Printf("圆的面积是%.2f,周长是%.2f", s, c)
}
