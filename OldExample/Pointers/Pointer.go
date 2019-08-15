// 指针

package Pointers

import (
	"fmt"
)

func main() {
	// 【&】获取值在内存中的地址
	// &T的数据类型是*T，*是指针，T是指针类型
	a := "指针测试"
	b := &a
	c := *b
	fmt.Println("test", c)

	var cat int = 1
	var str string = "banana"
	fmt.Printf("%v %p", &cat, &str)
	fmt.Println()

	v1 := 66
	v2 := 77
	swap(&v1, &v2)
	fmt.Println("指针修改后a,b为:", v1, v2)
}

// 通过指针修改值
func swap(a, b *int) {
	// 取a指针的值, 赋给临时变量t
	t := *a
	// 取b指针的值, 赋给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = t
}
