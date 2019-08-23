package main

import (
	"fmt"
)

func main() {
	s := Student{
		name:  "Lili",
		hobby: "sleep",
	}
	s.Work()
	Hello(s)
}

type Person interface {
	Name() string
	Career
}

type Career interface {
	Work()
}

type Student struct {
	name  string
	hobby string
}

func (s Student) Name() string {
	return s.name
}

func (s Student) Work() {
	fmt.Println("I am working!! " + s.Name())
}

// 使用接口作为参数
func Hello(person Person) {
	// 这里要将接口转换成具体的struct
	if obj, isTrue := person.(Student); isTrue {
		fmt.Println("转换成功, my hobby is " + obj.hobby)
		return
	}
	fmt.Println("Unknown type")
}
