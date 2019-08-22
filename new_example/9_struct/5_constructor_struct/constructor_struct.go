package main

import (
	"fmt"
	ptest "github.com/zhiwen-kooboo/GolangDemo/new_example/9_struct/package_for_test"
)

func main() {
	// 测试父struct Cat
	t := ptest.NewCat("zhangsan", "math teacher")
	fmt.Printf("%T\r\n", t)

	// 测试子struct BlackCat
	b := ptest.NewBlackCat("black1")
	fmt.Println(b.Color, b.Name)
	fmt.Printf("%T\r\n", b)
}
