package main

import (
	"fmt"
)

func main() {
	// struct的【==】【!=】比较
	a := role{1, "学生"}
	b := role{1, "学生"}
	c := role{1, "教师"}
	fmt.Println(a == b, a == c)
}

type role struct {
	id   int
	name string
}
