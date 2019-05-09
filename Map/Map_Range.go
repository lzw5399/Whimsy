package main

import (
	"fmt"
)

func main() {
	m := map[int]string{
		1: "test1",
		2: "test2"}

	for i, v := range m {
		fmt.Println(i, v)
	}

	//slice i是索引, v是值（拿到的是值的拷贝）,但是可以通过s[i]进行修改本身
	//map i是key, v是value(同样是拷贝)， 同样可以通过m[i]进行修改本身
	fmt.Println()
}
