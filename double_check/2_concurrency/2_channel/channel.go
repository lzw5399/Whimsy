package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 100)

	go func() {
		//time.Sleep(4 * time.Second)
		c <- 6
		close(c)
		fmt.Println("123")
	}()
	time.Sleep(time.Second * 1)

	select {
	case data, ok := <-c:
		fmt.Println(data, ok)
	default:
		fmt.Println("default")
	}
}
