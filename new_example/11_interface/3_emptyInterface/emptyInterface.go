package main

import (
	"fmt"
)

func main() {
	p := Person{
		Name: "michael",
	}
	TestEmptyInterface(p)
	TestEmptyInterfaceByTypeSwitch(p)
}

type Person struct {
	Name string
}

// 空接口作为参数，接收所有的type
func TestEmptyInterface(empty interface{}) {
	if obj, isTrue := empty.(Person); isTrue {
		fmt.Println("转换空接口成功, my hobby is " + obj.Name)
	}
}

// 通过type switch进行类型转换
func TestEmptyInterfaceByTypeSwitch(empty interface{}) {
	switch t := empty.(type) {
	case Person:
		fmt.Printf("type is %T", t)
	default:
		fmt.Println("未知类型")
	}
}
