package main

import "fmt"

func main() {
	var name string = "二狗子"
	var age = 18
	trueName := "李二狗"

	fmt.Printf("s1=%s s2=%d s3=%s \n", name, age, trueName)
	// 结果：s1=二狗子 s2=18 s3=李二狗

	// 变量交换
	var num1 int = 10
	var num2 int = 20
	num1, num2 = num2, num1
	fmt.Println(num1, num2)	// 20 10
}
