// 当前程序的包名
package main

// 导入其他的包
import (
	"fmt"
	_ "time"
)

// 全局常量的定义
const (
	Pi  = 3.14
	d = "1"
)

// 全局变量的定义
var (
	Name = "gopher"
)

// 一般类型声明
type (
	newType int
	newType2 string
)

// 结构体声明
type gopher struct { }

// 接口声明
type golang interface { }

func main() {
	var a int = 1
	var b uint = 1
	var c int8 = 1
	var d int16 = 1
	var e int32 = 1 // rune
	var f int64 = 1
	var g uint8 = 1 // byte
	var h uint16 = 1
	var i uint32 = 1
	var j uint64 = 1

	var k rune = 1
	var l byte = 1

	fmt.Println("hello from fmt2 ")
}
