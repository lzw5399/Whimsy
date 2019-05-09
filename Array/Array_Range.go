package main

import (
	"fmt"
)

func main() {
	team := [3]string{"hello", "sun", "tom"}

	for i, v := range team {
		if v == "sun" {
			team[i] = "666"
		}
		fmt.Println(i, team[i])
	}
}
