package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go hello()

	var input string
	fmt.Scanln(&input)
}

func hello() {
	for {
		fmt.Println("hello")
		time.Sleep(1 * time.Second)
	}
}
