package ch6

import "fmt"

func main() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b) // true
	fmt.Println(a == c) // false
	fmt.Println(b == c) // false

	//d := [3]int{1, 2}
	//fmt.Println(a == d) // invalid operation: a == d (mismatched types [2]int and [3]int)
}
