package main

import (
	"fmt"
)

func main() {
	// 直接初始化map的方式
	var m1 = map[rune]string{
		1: "233", 2: "444"}
	fmt.Println(m1)

	// 使用make初始化map
	m2 := make(map[string]string, 6)
	m2["test"] = "ok"
	fmt.Println(m2)

	// 赋值
	m2["k"] = "hello"

	// 判断是否存在
	if _, exist := m2["why"]; !exist {
		m2["why"] = "hi"
	}

	// 删除
	delete(m2, "test")
	fmt.Println(m2)

	// 对map进行foreach操作
	// 如果需要改变数组原值，使用map[k]进行操作
	for k, v := range m2 {
		fmt.Println("key", k, "value", v)
	}
}
