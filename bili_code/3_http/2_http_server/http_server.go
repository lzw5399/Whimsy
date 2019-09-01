package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 匹配主路由，并传一个处理函数
	http.HandleFunc("/", hello)
	// 启动服务并监听地址，传nil
	err := http.ListenAndServe("127.0.0.1:8880", nil)
	if err != nil {
		fmt.Println("failed: ", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		cost := time.Since(start)
		fmt.Println("cost:", cost.Seconds())
	}()
	fmt.Println("welcome to goland world!")
	fmt.Fprintf(w, "hello")
}
