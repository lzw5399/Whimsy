package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:80")
	fmt.Println(conn, err)
	conn2, err := net.Dial("tcp", "127.0.0.1:8080")
	fmt.Println(conn2)

	a, b := 1, 2
	a, b = b, a
	fmt.Println(a, b)
}
