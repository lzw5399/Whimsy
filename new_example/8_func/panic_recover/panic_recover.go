package main

import "fmt"

func main() {
	p1()
	p2()
	p3()
}

//
func p1() {
	fmt.Println("p1")
}

func p2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%v\r\n", err)
		}
	}()
	panic("this is panic")
}

func p3() {
	fmt.Println("p3")
}
