package main

import "fmt"

// 算术运算符
func main() {
	var num1 int = 4
	var num2 int = 2

	fmt.Println(num1 + num2) // 6
	fmt.Println(num1 - num2) // 2
	fmt.Println(num1 * num2) // 8
	fmt.Println(num1 / num2) // 2
	fmt.Println(num1 % num2) // 0
	num1++
	fmt.Println(num1) // 5
	num2--
	fmt.Println(num2) // 1
}
