package main

import "fmt"

// Goto 语句
func main() {
	var i int = 5

LOOP:
	for i < 10 {
		if i == 5 {
			i += 1
			goto LOOP
		}
		fmt.Printf("a is value: %d\n", i)
		i++
	}
}
