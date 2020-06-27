package ch6

import "fmt"

func main() {
	var a [3]int

	var b [3]int = [3]int{1, 2, 3}
	var c [3]int = [3]int{1, 2}

	q := [...]int{4, 5, 6}

	fmt.Println(a, b, c, q)
	// output: [0 0 0] [1 2 3] [1 2 0] [4 5 6]

	s := [...]string{"a", "b", 3: "c"}
	for k, v := range s {
		fmt.Println(k, v)
	}
	// output:
	//0 a
	//1 b
	//2
	//3 c

}
