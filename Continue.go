package main

import "fmt"

func main() {
label1:
	for i := 0; i < 3; i++ {
		for a := 1; a < 10; a++ {
			if a == 1 {
				fmt.Println(a)
				continue label1
			} else {
				fmt.Println("不是")
			}
		}
	}
}
