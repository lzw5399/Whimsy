package main

import "fmt"

func main() {
	// 初始容量0，元素个数0
	s := make([]int, 0)
	fmt.Println(s)
	pingPong(s)
	fmt.Printf("%p\r\n", s)
}

func pingPong(s []int) {
	s = append(s, 3)
	fmt.Printf("%p\r\n", s)
}
