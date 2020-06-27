package main

import "fmt"

func main() {

	t := "中"
	fmt.Printf(" UTF8 %x\n", t)
	fmt.Printf(" Unicode %U \n", t)

	s := "中华人民共和国"
	for _, c := range s {
		fmt.Printf("%[1]c %[1]x\n", c)
	}
}
