package main

import "fmt"

func main() {
	str := new(string)
	*str = "ninja"
	fmt.Println(*str)

	s := new([]int)
	*s = []int{1, 2, 3}
	for i := 0; i < len(*s); i++ {
		fmt.Println((*s)[i])
	}
}
