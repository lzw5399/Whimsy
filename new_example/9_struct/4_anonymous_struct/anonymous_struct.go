package main

import (
	"fmt"
)

func main() {
	// 定义一个匿名结构体，并初始化
	s2 := struct {
		id   int
		name string
	}{1, "233"}
	fmt.Println(s2.id, s2.name)

	// 匿名结构体的使用
	msg := &struct {
		name    string
		message string
	}{"lilei", "nice to meet you"}
	printMessage(msg)

	// 使用struct里包含【匿名struct】的情况
	p := Programmer{
		Name: "张三",
		Contact: struct {
			Phone, City string
		}{
			Phone: "6666",
			City:  "shanghai",
		},
	}
	fmt.Println(p)
}

func printMessage(obj *struct {
	name    string
	message string
}) {
	fmt.Printf("%T\r\n", obj)
	fmt.Println(obj.name, obj.message)
}

type Programmer struct {
	Name    string
	Contact struct {
		Phone, City string
	}
}
