package main

import (
	"fmt"
)

func main() {
	a := [5]int{}
	a[1] = 7
	fmt.Println(a)

	b := new([5]int)
	b[1] = 8
	fmt.Println(b)
}
