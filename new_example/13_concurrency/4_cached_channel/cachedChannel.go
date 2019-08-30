package main

import (
	"fmt"
	"sync"
)

func main() {
	// go 1.5以后就会默认执行下面这句
	// runtime.GOMAXPROCS(runtime.NumCPU())

	// 多个goroutine，主线程如何确保同时完成
	// 方案1、通过带缓存的chan来完成
	//c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go hello(c, i)
	}
	//for i := 0; i < 10; i++ {
	//	<-c
	//}

	// 方案2、通过sync.WithGroup来完成
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go hello2(&wg, i)
	}
	wg.Wait()
}

// 方案1、通过带缓存的chan来完成
func hello(c chan int, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	c <- index
}

// 方案2、通过sync.WithGroup来完成
func hello2(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()
}
