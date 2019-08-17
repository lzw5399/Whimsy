package main

import (
	"fmt"
)

func main() {
	m := make(map[int]map[int]int)
	m[0] = make(map[int]int)
	// m[0][9] = "test"

	value, isExist := m[0][9]
	fmt.Println(value, isExist)
}
