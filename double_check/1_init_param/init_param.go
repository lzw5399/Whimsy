package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:80")
	fmt.Println(conn, err)
	conn2, err := net.Dial("tcp", "127.0.0.1:8080")
	fmt.Println(conn2)

	a, b := 1, 2
	a, b = b, a
	fmt.Println(a, b)

	var by []byte = []byte{1, 2, 3, 4, 5}
	fmt.Println(len(by))

	time.Now()
}
