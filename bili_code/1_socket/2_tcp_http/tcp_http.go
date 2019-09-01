package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	// 新建一个tcp连接
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("error dialing", err)
		return
	}
	defer conn.Close()
	msg := "GET / HTTP/1.1\r\n"
	msg += "Host: www.baidu.com\r\n"
	msg += "Connection:close\r\n"
	msg += "\r\n\r\n"
	// 将msg写入到conn中
	n, err := io.WriteString(conn, msg)
	if err != nil {
		fmt.Println("write string failed", err)
		return
	}
	fmt.Println("send to baidu.com bytes: ", n)
	// 创建缓冲区
	buf := make([]byte, 4096)
	// 循环写入，并且在Read失败后进行break
	for {
		count, err := conn.Read(buf)
		fmt.Println("count:", count, " err:", err)
		if err != nil {
			break
		}
		// 打印结果
		fmt.Println(string(buf[0:count]))
	}
}
