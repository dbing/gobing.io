package main

import "fmt"

func main() {
	// 声明单个常量
	const pi = 3.1415926
	fmt.Printf("%T %v\n", pi, pi) // float64 3.1415926

	// 声明多个常量
	const (
		e = 2.7182818
		p = 3.1415926
	)
	fmt.Println(e, p) // 2.7182818 3.1415926
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Printf("%T a=%[1]d b=%d c=%d d=%d", a, a, b, c, d) // int a=1 b=1 c=2 d=2
}
