package main

import (
	"fmt"
)

func main() {
	// f1
	a, _ := f1(1, "2")
	fmt.Println("f1返回值", a)

	// f2
	c, d := f2("11", "222")
	fmt.Println("f2返回值", c, d)

	// f3
	fmt.Println("f3不定长度参数的个数为", f3(1, "1", "22", "dsds"))

	// 函数也是一个类型,用法一，直接指定外部函数
	e := f1
	e(1, "2")

	// 函数也是一个类型,用法二，匿名函数
	f := func(a string) {
		fmt.Println("当前参数是", a)
	}
	f("匿名函数！！")

	// 闭包
	g := f4(6)
	fmt.Println("闭包的输出值为", g(7))
}

// 一般形式的【输入参数】和【输出参数】
func f1(a int, b string) (string, int) {
	fmt.Println("hello", a, b)

	return "fn", 2
}

// 同类型参数的简写，返回值参数的提前命名
func f2(a, b string) (c string, d int) {
	fmt.Println("hello", a, b)
	c, d = "22", 3
	return
}

// 不定长度变参
func f3(a int, b ...string) (c int) {
	c = len(b)
	return
}

// 闭包， 输入【int】,返回【func(b int) int】
func f4(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
