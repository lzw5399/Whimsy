package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// 修改
	a := 1
	ModifyBasicType(&a)
	fmt.Println(a)

	// 修改结构体的
	s := TestStruct{
		A: "1",
		B: 2,
		C: "3",
	}
	ModifyStructPointer(&s)
	fmt.Println(s)
}

func ModifyBasicType(i interface{}) {
	v := reflect.ValueOf(i)

	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Int {
		panic("不是int类型")
	}

	v.Elem().SetInt(2)
}

func ModifyStructPointer(i interface{}) {
	v := reflect.ValueOf(i)

	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct || !v.Elem().CanSet() {
		panic("非法的数据类型或无法set")
	}

	v = v.Elem()

	// 修改所有的string的字段值
	for i := 0; i < v.NumField(); i++ {
		currentFiled := v.Field(i)
		if currentFiled.Kind() == reflect.String && currentFiled.CanSet() {
			currentFiled.SetString("修改后的值" + strconv.Itoa(i))
		}
	}
}

type TestStruct struct {
	A string
	B int
	C string
}
