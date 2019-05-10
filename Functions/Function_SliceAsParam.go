package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2}
	C(a)
	fmt.Println(a)
}

func C(s []int) {
	s[0] = 666
	s[1] = 777
	fmt.Println(s)
}
