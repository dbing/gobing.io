package main

import "fmt"

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // 255 0 1
	fmt.Printf("%08b %08b\n", 255, 0)

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // 127 -128 1
	fmt.Printf("%08b %08b\n", 127, -128)

}
