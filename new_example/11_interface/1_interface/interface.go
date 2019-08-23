package main

import (
	"fmt"
)

func main() {
	var a USB
	a = PhoneConnector{
		name: "hello",
	}
	a.Connect()
	DisConnect(a)
}

// 定义接口
type USB interface {
	Name() string
	Connect()
}

// 定义struct
type PhoneConnector struct {
	name string
}

// 定义method让struct实现interface
func (p PhoneConnector) Name() string {
	return p.name
}

func (p PhoneConnector) Connect() {
	fmt.Println("Connect: " + p.name)
}

func DisConnect(usb USB){
	fmt.Println("DisConnected")
}
