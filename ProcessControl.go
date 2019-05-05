// 流程控制
package main

import (
	"fmt"
	_ "fmt"
	"strconv"
	_ "strconv"
)

func main() {
	// if else分支语句
	/*if b := "aa"; len(b) == 1 {

	} else {
		fmt.Println("特殊写法,b的长度为" + strconv.Itoa(len(b)))
	}*/

	// for循环
	/*for i := 0; i < 10; i++ {
		fmt.Println(i)
	}*/

	// 省略初始语句的for循环
	step := 1
	for ; step <= 5; step++ {
		fmt.Println(step)
	}

	// 无限循环写法一
	var i int
	for ; ; i++ {
		if i > 10 {
			break
		}
	}
	fmt.Println("i的最终值为" + strconv.Itoa(i))

	// 无限循环写法二
	for {
		fmt.Println("3")
	}

	// 相当于while(k<=10)
	var k int
	for k <= 10 {
		k++
	}
}
