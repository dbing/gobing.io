package main

import "fmt"

func main() {
	const (
		num1 = iota
		num2
		num3 = 30
		num4
	)
	fmt.Println(num1, num2, num3, num4) // 0 1 30 30

	// 定义类型
	type Weekday int

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	// Output: 0 1 2 3 4 5 6

	const (
		_   = 1 << (10 * iota)
		KiB  // 1024
		MiB  // 1048576
		GiB  // 1073741824
		TiB  // 1099511627776
		PiB  // 1125899906842624
		EiB  // 1152921504606846976
		//ZiB
		//YiB
	)
	fmt.Println(KiB, MiB, GiB, TiB, PiB, EiB)
}
