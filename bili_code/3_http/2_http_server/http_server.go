package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 匹配路由，并传一个处理函数
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/user/login", handleLogin)

	// 启动服务并监听地址，传nil
	err := http.ListenAndServe("127.0.0.1:8880", nil)
	if err != nil {
		fmt.Println("failed: ", err)
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	//start := time.Now()
	//defer func() {
	//	cost := time.Since(start)
	//	fmt.Println("cost:", cost.Seconds())
	//}()
	fmt.Println("welcome to goland world!")
	// 这个是向网页输出的内容
	fmt.Fprintf(w, "hello")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request login")
	// 这个是向网页输出的内容
	fmt.Fprintf(w, "login")
}
