// gobing.io/ch2/variable.go
package main

import "fmt"

func main() {
	var num1, num2, num3 int = 1, 2, 3
	// 写法二：一次声明多个变量并赋值（不同类型）
	var (
		a string = "Hello"
		b int    = 20
		c bool   = true
	)
	// 写法三：使用类型自推导写法（省略类型）
	var num4, str5 = 4, "55"
	// 写法四：多个变量定义（简短写法）
	num6, num7 := 6, 7
	fmt.Println(num1, num2, num3)					// 1 2 3
	fmt.Printf("a 类型:%T 值：%s\n", a, a)		// a 类型:string 值：Hello
	fmt.Printf("b 类型:%T 值：%d \n", b, b)	// b 类型:int 值：20
	fmt.Printf("c 类型:%T 值：%t \n", c, c)	// c 类型:bool 值：true
	fmt.Printf("num4 类型:%T 值:%d \n", num4, num4) // num4 类型:int 值:4
	fmt.Printf("str5 类型:%T 值:%v \n", str5, str5) // str5 类型:string 值:55
	fmt.Println(num6, num7)							// 6 7

}
