package main

import "fmt"

// 指针
func main() {
	// 指定指针类型
	var m int8 = 100
	var n *int8 = &m
	fmt.Printf("%v %v %v %v\n", m, &m, &n, *n)
	*n = 10
	fmt.Println(m, *n, n)	// 10 10 0xc00001c09a

	// 自动推断指针类型
	x := 1
	p := &x
	fmt.Printf("%T %[1]v %v \n", p, *p) // *int 0xc0000b4008 1

	*p = 2
	fmt.Println(x, &x)

	fmt.Println(&x == p)

	// 指针比较
	var y, z int
	fmt.Println(&y == &y, &y == &z, &y == nil) // true false false
	// 指向同一个变量，或者

	// 模拟 ++v 操作
	v := 1
	fmt.Println(incr(&v))
	fmt.Println(incr(&v))
}

func incr(p *int) int {
	*p++
	return *p
}
