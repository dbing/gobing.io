package main

import "fmt"

func main() {
	var b1 bool = true
	var b2 bool = false

	fmt.Println(b1 && b2) // true
	fmt.Println(b1 || b2) // false
	fmt.Println(!b1)      // false
}
