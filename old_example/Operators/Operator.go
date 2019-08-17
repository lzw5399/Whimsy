// 运算符
package Operators

import "fmt"

const ()

func main() {
	a := 6 & 11  // 按位与
	b := 6 | 11  // 按位或
	c := 6 ^ 11  // 按位异或
	d := 6 &^ 11 // 将运算符左边数据相异的位保留，相同位清零
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(1 << 10)
	fmt.Println(1024 >> 10)

	if a != b {
		fmt.Println(a != b)
	}

	fmt.Println("done")
}
