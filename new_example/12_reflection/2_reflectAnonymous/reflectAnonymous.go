package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := Monkey{
		Animal: Animal{
			Id:   122,
			Name: "Tom",
			Age:  5,
		},
		Skill: "tree",
	}
	t := reflect.TypeOf(m)

	// 取整个匿名字段的内容
	// 真要遍历的时候可以这么用
	// t.Field(0).Type.Kind() == reflect.Struct
	fmt.Printf("%#v\r\n", t.Field(0))
	// 取匿名字段中，某个成员的内容,
	// 第一个0代表取的是第一个字段，即【struct】
	// 第二个0代表取【struct】中第一个字段
	fmt.Printf("%#v\r\n", t.FieldByIndex([]int{0, 1}))
}

type Animal struct {
	Id   int
	Name string
	Age  int
}

type Monkey struct {
	Animal
	Skill string
}
