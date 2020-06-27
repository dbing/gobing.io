package main

import "fmt"

func main() {
	// 函数可以没有参数或者有多个参数（但不支持有默认值）
	fmt.Println(add(1, 2))

	// 两个或多个连续的参数类型相同，除最后一个外，前面的类型声明可以省略
	fmt.Println(sum(6, 6))

	// 多返回值
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// 命名函数返回值（使其具有意义，可作文档使用）
	x, y := split(17)
	fmt.Println(x, y)

	// 可变参数

	// 匿名函数
}

func add(x int, y int) int {
	return x + y
}

func sum(x, y int) int {
	return x + y
}

// 多返回值
func swap(x, y string) (string, string) {
	return y, x
}

// 没有参数的 return 语句返回结果的当前值 x y
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
