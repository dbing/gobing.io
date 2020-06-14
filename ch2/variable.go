package main

import "fmt"

func main() {
	// 定义变量 s1 和 s2
	var s1 string
	var s2 string = "hello"
	s3 := "Hi"
	fmt.Printf("s1=%s s2=%s s3=%s \n", s1, s2, s3)

	// 同时定义变量
	var num1, num2, num3 = 1, 2, 3
	var num4, s5 = 4, "555"
	num6, num7 := 6, 7

	fmt.Println(num1, num2, num3)
	fmt.Printf("num4=%d; s5 type=%T value=%s \n", num4, s5, s5)
	fmt.Println(num6, num7)
}
