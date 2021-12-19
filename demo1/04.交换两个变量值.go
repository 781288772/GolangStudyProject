package main

import "fmt"

func main() {
	var num1, num2, temp int
	num1 = 10
	num2 = 20
	temp = num1
	num1 = num2
	num2 = temp

	fmt.Println(num1, num2)

}
