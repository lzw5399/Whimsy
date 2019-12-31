// 常量&枚举
package main

import "fmt"

// 星期枚举例子
const (
	a      = 'A'
	Monday = iota
	Tuesday
	Wednesday = 5
	Thursday = iota
	Friday
	Saturday
	Sunday
)

func main() {
	var x byte = 'a'
	fmt.Println(typeof(x))
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
	fmt.Println(Sunday)
}

// 获取当前变量的类型
func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
