// 数组
package Array

import (
	"fmt"
)

func main() {
	// 全部赋值
	a := [2]int{1, 6}
	var b [2]int
	b = a
	fmt.Println(b)

	// 部分赋值
	c := [3]int{1}
	fmt.Println(c)

	// 索引部分赋值 & 元素推测
	d := [...]int{1, 3: 9, 4}
	fmt.Println(d)

	// 指向数组的指针【是个指针】
	var e *[5]int = &d
	fmt.Println(e)

	p := new([10]int)
	fmt.Println("p", p)

	// 指针数组【是个数组】
	f, g := 1, 2
	h := [2]*int{&f, &g}
	fmt.Println(h)
}
