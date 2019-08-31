package main

import (
	"fmt"
	"time"
)

func main() {
	// getChannelValueBySelect()

	setChannelValueBySelect()

	// timeoutSample()
}

// 使用select给获取channel的值
func getChannelValueBySelect() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hell0"
	// close(c1)
	close(c2)

	<-o
}

// 使用select给channel赋值
func setChannelValueBySelect() {
	c := make(chan int)
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()
	for {
		select {
		case c <- 0:
		case c <- 1:
		case c <- 2:
		}
	}
}

// select超时例子
func timeoutSample() {
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case c2 := <-time.After(3 * time.Second):
		fmt.Println("timeout ttttt", c2)
	}
}
