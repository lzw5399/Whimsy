package main

import (
	"fmt"
)

func main() {
	// 从新声明一个slice
	// var a []int

	// 从数组里面，进行切片生成
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[1:len(a)]
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// 使用make声明
	s1 := make([]int, 3, 9)
	fmt.Println(len(s1), cap(s1))
}
