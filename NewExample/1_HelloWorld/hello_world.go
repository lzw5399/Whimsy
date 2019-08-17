package main

import (
	test "fmt"
	_ "io"
)

// 		常量
const PI = 3.14
const (
	A = 1
	b = 2
)

// 全局变量
var class = "golang"
var (
	c = "1"
	D = 3
)

// 一般type
type newType rune
type (
	e string
	f uint
)

// 定义type
type gopher struct{}

// 定义接口
type golang interface{}

func main() {
	test.Println(PI)

	test.Println(class)

	test.Println(c + string(D))
}
