package main

import "fmt"

func main() {
	// 写法一：求 1 到 10 数字之和
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 写法二：可以让前置、后置语句为空
	//total := 1
	//for ; total < 3; {
	//	fmt.Println("loop:", total)
	//	total += total
	//}
	//fmt.Println("end:",total)

	// 写法三：省略分号，Go 实现 while 写法
	total := 1
	for total < 3 {
		fmt.Println("loop:", total)
		total += total
	}
	fmt.Println("end:", total)

	// 死循环
	for {
		fmt.Println("loop")
	}
}
