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
	// 整形
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

	// 别名整形
	var k rune = 1 // int32
	var l byte = 1 // uint8

	// 浮点型
	var m float32 = .1
	var n float64 = 2.2

	// 复数 z=a+bi（a,b均为实数）的数称为复数，其中a称为实部，b称为虚部，i称为虚数单位
	var o complex64 = 1.11
	var p complex128 = 2.22

	fmt.Println("hello from fmt2 ")
}
