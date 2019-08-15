package main

import (
	"fmt"
)

func main() {
	// 方式一
	/*var m map[int]string // 中括号中为key，外面为value
	m = map[int]string{
		1: "dsd",
		2: "dsdsd"}*/

	// 方式二
	a := make(map[int]string, 3)
	a = map[int]string{
		1: "你好",
		2: "hello"}
	a[3] = "ok"
	delete(a, 2)
	fmt.Println(a)
}
