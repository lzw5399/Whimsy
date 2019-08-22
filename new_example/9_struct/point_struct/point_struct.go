package main

import (
	"fmt"
)

func main() {
	// struct直接作为参数和返回值
	p := chen(Point{2, 3}, 4)
	fmt.Println(p.X, p.Y)

	// struct指针 作为参数和返回值
	p2 := chen2(&Point{3, 4}, 5)
	p3 := *p2
	fmt.Println(p2.X, p2.Y, p3.X, p3.Y)

	// 测试,得出结论：
	// 1、func直接接收struct接收的是【对象的拷贝】
	// 2、func接收struct指针，接收的是【对象的引用】
	a := Point{1, 2}
	b := &Point{1, 2}
	changeValue(a)
	changePointValue(b)
	fmt.Println("直接使用struct，并改变值后的对象值为", a.X, a.Y)
	fmt.Println("使用struct指针，并改变值后的对象值为", b.X, b.Y)
}

type Point struct {
	X, Y int
}

// struct直接作为参数和返回值
func chen(p Point, bit int) Point {
	return Point{
		X: p.X * bit,
		Y: p.Y * bit,
	}
}

// struct指针 作为参数和返回值
func chen2(p *Point, bit int) *Point {
	fmt.Println("函数内struct指针的值", p.X, p.Y)
	return &Point{
		X: p.X * bit,
		Y: p.Y * bit,
	}
}

func changeValue(p Point) {
	p.X, p.Y = 222, 333
}

func changePointValue(p *Point) {
	p.X, p.Y = 222, 333
}

// 这是构造函数啊，输入值返回一个 对应struct的指针
func NewPoint(x int, y int) *Point {
	return &Point{X: x, Y: y}
}
