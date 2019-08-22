package main

import (
	"fmt"
)

func main() {
	// defer普通调用
	//fmt.Println("1")
	//f11()
	//fmt.Println("2")

	// defer闭包调用
	f22()
}

// defer普通调用
func f11() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	fmt.Println("d")
}

// defer的闭包调用
func f22() {
	for a := 0; a < 3; a++ {
		defer func() { fmt.Println(a) }()
	}
}
