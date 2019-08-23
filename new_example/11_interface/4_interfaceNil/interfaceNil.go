package main

import (
	"fmt"
)

// 只有当接口存储的类型和对象都为nil时，接口才等于nil，所以为true
func main() {
	// 空指针，所以为nil
	var a interface{}
	fmt.Println(a == nil)

	// 虽然b对象为nil，但是a接口指向的对象不为nil，所以为false
	var b *int = nil
	a = b
	fmt.Println(a == nil)
}
