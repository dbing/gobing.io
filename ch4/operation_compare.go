package main

import "fmt"

// 比较运算
func main() {

	// 相同类型比较运算
	var n1 int = 1
	var n2 int = 2

	fmt.Println(n1 == n2) // false
	fmt.Println(n1 != n2) // true
	fmt.Println(n1 < n2)  // true
	fmt.Println(n1 <= n2) // true
	fmt.Println(n1 > n2)  // false
	fmt.Println(n1 >= n2) // false

	// 不同类型不能比较
	var num1 int16 = 1
	var num2 int16 = 2
	//var num3 int32 = 3
	fmt.Println(num1 > num2) // false
	//fmt.Println(num1 > num3) // mismatched types int16 and int32
}
