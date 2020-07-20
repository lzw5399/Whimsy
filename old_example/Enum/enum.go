package main

import "fmt"

const (
	a = "A"
	b = 8
	c = iota //2
	d        //3
)

func main() {
	fmt.Println(a, b, c, d)
}
