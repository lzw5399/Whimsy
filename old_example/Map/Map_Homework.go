package main

import (
	"fmt"
)

func main() {
	m1 := map[int]string{
		0: "j",
		1: "a",
		2: "b",
		3: "c"}

	m2 := make(map[string]int, len(m1))
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println("m1", m1)
	fmt.Println("m2", m2)
}
