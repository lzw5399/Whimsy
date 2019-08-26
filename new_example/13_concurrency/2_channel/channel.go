package main

import (
	"fmt"
)

func main() {
	// 创建一个channel对象，里面存的数据是bool
	c := make(chan int)

	go func() {
		fmt.Println("run run run")
		// 将true存进channel, 这里的关键是因为我们定义了一个chan，并且在【主线程】外取chan的值
		// 所以goroutine里面【主线程】会等待chan被赋值，再所以【gogogo】会被稳定打印
		c <- 1
	}()
	x, y := <-c // 取出channel
	close(c)
	fmt.Println(x, y)
}
