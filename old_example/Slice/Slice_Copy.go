package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9}
	fmt.Printf("copy之前的s1地址:%p,长度:%v,容量%v\n", s1, len(s1), cap(s1))
	fmt.Printf("copy之前的s2地址:%p,长度:%v,容量%v\n", s2, len(s2), cap(s2))

	copy(s2, s1) //将s1赋值给s2
	fmt.Printf("copy之前的s1地址:%p,长度:%v,容量%v\n", s1, len(s1), cap(s1))
	fmt.Printf("copy之前的s2地址:%p,长度:%v,容量%v\n", s2, len(s2), cap(s2))
	fmt.Println("当前s2", s2)
}
