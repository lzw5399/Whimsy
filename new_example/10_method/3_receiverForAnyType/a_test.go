package a_test

import (
	"fmt"
	"testing"
)

func TestReceiver(t *testing.T) {
	var b myInt
	fmt.Println(b.IsZero())

	b.Add(6)
	fmt.Println(b)
}

// 为基本类型添加接收器，先新建一个type来指向基本类型
// 然后使用这个新建的type来作为接收器
type myInt int

func (i myInt) IsZero() bool {
	return i == 0
}

// 因为int是值类型，所以这边需要指针来传引用
func (i *myInt) Add(value int) {
	*i += myInt(value)
}
