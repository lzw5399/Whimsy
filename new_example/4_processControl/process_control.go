package main

import (
	"fmt"
)

func main() {
	// if
	a := 1
	if a := 2; a < 3 {
		fmt.Println(a)
	}
	fmt.Println(a)

	// for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	var j int
	for j <= 11 {
		fmt.Println(j)
		j++
	}

	// switch
	// 使用常量进行判断【并有初始化表达式】
	switch a := 1; a {
	case 2:
		fmt.Println("1")
	case 3:
		fmt.Println("2")
	default:
		fmt.Println("3")
	}

	// 使用表达式进行判断【无初始化表达式】
	xx := 1
	switch {
	case xx > 1:
		fmt.Println("1")
	case xx == 1:
		fmt.Println("2")
	default:
		fmt.Println("3")
	}
}
