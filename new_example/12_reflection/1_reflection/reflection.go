package main

import (
	"fmt"
	"reflect"
)

func main() {
	u := User{
		Id:   11,
		Name: "James",
		Age:  22,
	}
	u.Hello()
	Info(u)

	var a rune
	Info(a)
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("hello world")
}

func Info(obj interface{}) {
	t := reflect.TypeOf(obj) // 对象的类型信息
	fmt.Println(t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("当前对象类型不为struct")
		return
	}

	v := reflect.ValueOf(obj) // 对象的值信息

	// 打印字段信息
	for i := 0; i < t.NumField(); i++ {
		tt := t.Field(i) // 字段的类型信息

		value := v.Field(i).Interface() // 字段的值信息
		fmt.Printf("%s:%v=%v\r\n  ", tt.Type, tt.Name, value)
	}

	// 打印方法信息
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%s: %v\r\n", m.Name, m.Type)
	}

}
