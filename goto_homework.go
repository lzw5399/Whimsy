package main

import "fmt"

func main() {
label1:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			continue label1
		}
	}
}
