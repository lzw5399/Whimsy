package main

import "fmt"

func main() {
	// 创建一个channel对象，里面存的数据是bool
	c := make(chan [2]int)

	go func() {
		fmt.Println("run run run")
		// 将true存进channel, 这里的关键是因为我们定义了一个chan，并且在【主线程】外取chan的值
		// 所以goroutine里面【主线程】会等待chan被赋值，再所以【gogogo】会被稳定打印
		c <- [2]int{22, 33}
		close(c)
	}()

	// range chan只有一个返回值
	// 在迭代的过程中一直在等待chan，有一个值进去，才会打印
	for v := range c {
		fmt.Println(v)
	}
}
