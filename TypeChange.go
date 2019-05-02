package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*var a uint32 = 1  //多种整形格式之间也必须要强制转换
	var b uint = 2
	var c int8 = 3
	a = uint32(c)
	b = uint(a)


	d := uint8(a) + uint8(b) + uint8(c)*/
	var a int = 65
	d := strconv.Itoa(a)
	fmt.Println(d)
}
