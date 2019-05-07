package main

import (
	"fmt"
)

func main() {
	team := [3]string{"hello", "sun", "tom"}

	for k, v := range team {
		fmt.Println(k, v)
	}
}
