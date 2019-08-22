package main

import (
	"fmt"
)

func main() {
	// 第一种使用struct的方法
	p1 := new(person)
	p1.id = 1
	p1.Name = "zhangsan"
	fmt.Printf("%T\r\n",p1)

	// 第二种，也是最常用，直接指定属性名称来指定值
	p2 := person{Name: "233", id: 2333}

	// 第三种，通过参数的顺序，指定成员的值
	p3 := person{23231, "233"}

	// struct也是一种类型
	aa := person{}
	fmt.Println(aa) // 这边输出类型里面属性的零值

	fmt.Println(p1.id, p2.Name, p3.id)
}

type person struct {
	id   rune
	Name string
}


