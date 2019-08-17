package main

import (
	"fmt"
)

func main() {
	test := D
	a := 1
	test(&a)
	fmt.Println(a)
}

func D(p *int) {
	*p = 66
	fmt.Println(*p)
}
