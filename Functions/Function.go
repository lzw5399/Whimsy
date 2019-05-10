package main

import (
	"fmt"
)

func main() {
	a := A(2, 3, 4, 5, 6)
	fmt.Println(a)
}

// func 函数名(参数A int,参数B string)(int, string)
// 第二个括号为返回值，如果无返回值可以去掉，一个返回值可不要括号
func A(p1 ...int) (a string) {
	fmt.Printf("%T\n", p1)
	fmt.Println(p1)
	a = "121212"
	return a
}
