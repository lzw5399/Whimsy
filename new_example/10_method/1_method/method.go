package main

import "fmt"

func main() {
	a := &A{
		Number: 0,
	}
	// 调用接收器
	fmt.Println(a.IsZero())
}

type A struct {
	Number int
}
type B struct {
	Number int
}

// 接收器，类似C#中的拓展方法
func (a A) IsZero() bool {
	return a.Number == 0
}

func (b B) IsZero() bool {
	return b.Number == 0
}
