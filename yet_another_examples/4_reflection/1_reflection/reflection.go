package main

import (
	"fmt"
	"reflect"
)

// scope
// 1.打印出类型的名字
// 2.打印出是什么类型
// 3.打印出所有的字段
// 4.打印方法信息
func main() {
	obj := Test{name: "名字而已"}
	Info(obj)
}

type Test struct {
	name string
}

func (t Test) Hi() {
	fmt.Println("Hi")
}

func Info(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println(t.Name())

	if t.Kind() != reflect.Struct {
		fmt.Println("kind不是结构体")
	}

	// 打印所有字符串的字段
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type.Kind() == reflect.String {
			fmt.Println("是字符串", t.Field(i))
		}
	}

	// 打印方法名称
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println("方法有", t.Method(i).Name)
	}
}
