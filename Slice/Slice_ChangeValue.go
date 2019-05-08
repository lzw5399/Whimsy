package main

import (
	"fmt"
)

func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sa := a[2:5]
	sb := a[3:5]
	sb[0] = 999

	fmt.Println("sa", sa)
	fmt.Println("sb0", sb[0])
}
