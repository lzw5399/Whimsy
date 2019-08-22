package main

import (
	"fmt"
)

func main() {
	p1 := new(person)
	p1.id = 1
	p1.name = "zhangsan"

	p2 := person{name: "233", id: 2333}

	p3 := person{23231, "233"}

	fmt.Println(p1.id, p2.name, p3.id)

	aa := person{}
	fmt.Println(aa)
}

type person struct {
	id   rune
	name string
}
