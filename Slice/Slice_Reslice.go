package main

import (
	"fmt"
)

func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sa := a[2:5]
	fmt.Println("sa", len(sa), cap(sa))

	sb1 := sa[0:]
	fmt.Println("sb1", sb1)

	sb2 := sa[4:8]
	fmt.Println("sb2", sb2)
}
