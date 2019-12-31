package main

import (
	"fmt"
)

func main() {
	num := "233"
	if l := len(num); l > 2 {
		fmt.Println("L大于2")
	}
}
