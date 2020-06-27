package main

import (
	"fmt"
	"time"
)

func main() {
	// 写法一：有 switch 条件
	var grade string = "A"
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	case "C":
		fmt.Println("及格")
	default:
		fmt.Println("很尴尬，不及格")
	}

	// 写法二：没有条件的 switch 同 `switch true` 一样。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// fallthrough 无条件的继续执行下面所有 case （穿透）
	m := 2
	switch m {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
		fallthrough
	case 3:
		fmt.Println("case 3")
	}
}
