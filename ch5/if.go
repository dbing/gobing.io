package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 1 || args == nil {
		return
	}
	score, e := strconv.Atoi(args[1])
	if e != nil {
		fmt.Println(e)
		return
	}
	if score > 100 || score < 0 {
		fmt.Println("输入分值有误")
		return
	}
	if score >= 60 {
		fmt.Println("恭喜：及格了")
	}
	if score < 60 {
		fmt.Println("很遗憾：未通过考核")
	}
}
