package main

import (
	"fmt"
)

func main() {
	a, b := []int{1}, []int{2}
	B(a, b)
	fmt.Println(a, b)
}

func B(s ...[]int) {
	s[0] = []int{222}
	s[1] = []int{333}
	fmt.Println(s)
}
