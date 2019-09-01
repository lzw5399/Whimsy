package main

import (
	"fmt"
	"net" // socket编程需要的包
)

func main() {
	fmt.Println("start server...")
	// 监听端口, tcp模式
	listen, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil {
		fmt.Println("error occurred", err)
		return
	}
	// 无线循环，接受端口的返回值
	// 并且开启goroutine处理
	for {
		fmt.Println("1")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		// 建立缓冲区
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read failed", err)
			return
		}
		fmt.Println("read:", string(buf))
	}
}
