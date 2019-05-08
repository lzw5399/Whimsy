// 多维数组
package Array

import (
	"fmt"
)

func main() {
	a := [2][3]int{
		{1, 1, 1},
		{2, 2, 2}}
	c := [2][3]int{
		{1, 1, 1},
		{2, 2, 2}}
	fmt.Println(a == c)

	// b := &a
	fmt.Printf("%T", a)
}
