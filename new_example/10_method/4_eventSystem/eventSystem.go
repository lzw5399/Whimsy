package main

import (
	"fmt"
)

func main() {
	// 无论是普通函数还是结构体的方法，只要它们的签名一致，与它们签名一致的函数变量就可以保存普通函数或是结构体方法
	var de func(value string)
	p := person{}
	de = p.do
	de("Hi")

	de = funcDo
	de("Hi")
}

type person struct{}

func (p *person) do(value string) {
	fmt.Println("接收器方法: " + value)
}

func funcDo(value string) {
	fmt.Println("一般函数: " + value)
}
