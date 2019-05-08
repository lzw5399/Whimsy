// Go中的默认值
// Go中类型的声明
package Variable_Const

import (
	"fmt"
	_ "math"
)

/*type(
	整形 int
)*/

func main() {
	// var param [5]int   // 数组，默认值为[0 0 0 0 0]
	//var param 整形      // 类型别名

	// param := 1.12
	// fmt.Println(typeof(param))  //输出变量的type

	var a, b, c int   // 先并行声明
	a, b, c = 1, 2, 3 // 再并行赋值

	var e, f, g = 1, 3, 4 // 并行声明且赋值

	h, i := 5, 6 // 省略var的 并行声明且赋值
	fmt.Println(a + b + c + e + f + g + h + i)
}

// 获取当前变量的类型
/*func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}*/
