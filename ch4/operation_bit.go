package main

import "fmt"

// 位运算符
func main() {
	var x uint8 = 1
	var y uint8 = 255
	fmt.Printf("%08b %d\n", x, x)    // 00000001 1
	fmt.Printf("%08b %[1]d\n", y)    // 11111111 255
	fmt.Printf("%08b %[1]d\n", x&y)  // 00000001 1
	fmt.Printf("%08b %[1]d\n", x|y)  // 11111111 255
	fmt.Printf("%08b %[1]d\n", x^y)  // 11111110 254
	fmt.Printf("%08b %[1]d\n", x&^y) // 00000000 0
	fmt.Printf("%08b %[1]d\n", x<<2) // 00000100 4
	fmt.Printf("%08b %[1]d\n", y>>2) // 00111111 63
}

//
