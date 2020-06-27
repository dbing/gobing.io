package main

import "fmt"

func main() {
	var apples int32 = 1
	var oranges int16 = 2
	//var compote int = apples + oranges 	// (mismatched types int32 and int16)
	var compote int = int(apples) + int(oranges)
	fmt.Println(compote) // 3

	f := 1.56
	i := int(f)
	fmt.Println(f, i)

	f = 1.99
	fmt.Println(int(f))
}
