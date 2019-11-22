package main

import (
	"fmt"
)

func main() {
	a := 1 << 10 // 左移，乘以2的n次方
	fmt.Println(a)

	fmt.Println(^2) // 未知是干嘛的

	d := 6 & 11  //位运算符之 【按位与运算符】
	e := 6 | 11  //位运算符之 【按位或运算符】
	f := 6 ^ 11  //位运算符之 【按位异或运算符】
	g := 6 &^ 11 //位运算符之 【按位零运算符】
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
