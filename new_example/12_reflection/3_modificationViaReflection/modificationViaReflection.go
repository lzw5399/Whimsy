package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 基本类型的反射修改
	x := 123
	SetBasicTypeValue(&x)
	fmt.Println(x)

	// struct的 反射修改
	l := lion{
		Id:   100,
		Name: "Xinba",
		Age:  99,
	}
	SetObjFields(&l)
	fmt.Println(l)

	// l.Ao("ping")
	// 由于只是调用方法，不是修改原对象的值，遂这个指针，可传可不传，当然最好传，这里偷个懒
	InvokeMethodViaReflect(l)
}

type lion struct {
	Id   int
	Name string
	Age  int
}

func (l lion) Ao(str string) string {
	if str == "ping" {
		fmt.Println("pong")
		return "这是返回值: pong"
	}
	fmt.Println("unknown")
	return "这是返回值: unknown"
}

// 反射修改基本类型的值
func SetBasicTypeValue(obj interface{}) {
	// 这里不做过多的类型判断了
	v := reflect.ValueOf(obj) // 想要修改原对象信息，这里必须传一个指针
	v.Elem().SetInt(555)
}

// 反射修改结构体里的值
func SetObjFields(obj interface{}) {
	v := reflect.ValueOf(obj)

	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct || !v.Elem().CanSet() {
		fmt.Println("无效数据结构或不能set")
		return
	}
	// 因为【ValueOf】和【Elem】都是【reflect.Value】类型的
	// 这里确定了obj是pointer，所以可以直接这么用，获取原来对象的值
	v = v.Elem()

	// 遍历struct的field
	for i := 0; i < v.NumField(); i++ {
		currentField := v.Field(i)
		switch currentField.Kind() {
		case reflect.String:
			currentField.SetString("哈哈")
		case reflect.Int:
			currentField.SetInt(int64(i))
		}
	}
}

// 封装方法还是需要传指针，然后里面类型判断
// 这里不做判断了，只展示最核心的部分
func InvokeMethodViaReflect(obj interface{}) {
	v := reflect.ValueOf(obj)
	// 获取方法
	method := v.MethodByName("Ao")
	// 准备传入的参数
	args := []reflect.Value{reflect.ValueOf("ping")}
	// 调用方法，并获取返回值
	result := method.Call(args)
	fmt.Println("调用返回值结果=>", result[0])
}
