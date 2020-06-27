package main

import "fmt"

func main() {
	p := new(int)
	fmt.Printf("%T %v \n", p, p) // *int 0xc0000b4008
	fmt.Println(*p)	// 0

	*p = 2
	fmt.Println(*p)	// 2

}
