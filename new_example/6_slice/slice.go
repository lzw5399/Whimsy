package main

import (
	"fmt"
)

func main() {
	// 定义一个空的slice
	var s1 []rune
	fmt.Println(s1)

	// 从array中定义一个slice
	a1 := []rune{1, 2, 3, 4, 5}
	s2 := a1[:] // 这里索引包含【0】，不包含【3】
	fmt.Println(s2)

	// 使用make，从新创建一个slice
	s3 := make([]rune, 6, 10)
	fmt.Println(s3, len(s3), cap(s3))

	// reslice
	a2 := [10]rune{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ss1 := a2[1:3]
	ss2 := ss1[0:2]
	fmt.Println("reslice不超过len的", ss2)
	ss3 := ss1[2:9]
	fmt.Println("reslice超过len长度", ss3)

	// slice的append
	sss1 := make([]rune, 5, 10)
	sss2 := make([]rune, 1, 2)
	sss1 = append(sss1, 1, 2, 3, 4, 5, 6)  //
	sss1 = append(sss1, sss2...)
	fmt.Println("slice的append", sss1, cap(sss1))
}
