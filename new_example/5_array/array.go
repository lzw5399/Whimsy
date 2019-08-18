package main

import (
	"fmt"
)

func main() {
	// 数组全定义
	var a1 [2]int = [2]int{1, 5}
	fmt.Println(a1)

	// 数组简化定义
	a2 := [2]string{"test"} // string是空字符串
	fmt.Println(a2)

	// 指向数组的指针
	a3 := [...]rune{4: 222}
	p := &a3
	fmt.Println(p)

	// 指向数组的指针，通过new()创建
	p2 := new([6]rune)
	p2[3] = 66666
	fmt.Println(p2)

	// 指针数组
	x, y := 1, 2
	a4 := [2]*int{&x, &y}
	fmt.Println(a4)

	// 数组的比较
	a5 := [2]int{1, 2}
	a6 := [...]int{1, 2}
	fmt.Println(a5 == a6) // 返回true
	fmt.Println(cap(a6))
}
