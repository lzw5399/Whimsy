package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[:]
	fmt.Println("当前s2", s2)
}
