package main

import (
	"fmt"
	"math" // math包里面包含一些整形的最大值、最小值常量
	"strconv"
)

type (
	字符串 string // 设置string类型别名为【字符串】
)

func main() {
	var (
		a          [11]int          // 变量的默认值
		str        字符串     = "类型别名" // 类型别名
		q, w, e, r         = 0, 1, 2, 3
	)

	t := "t"

	var (
		ff = 122
	)

	fmt.Println(a, ff)
	fmt.Println(str)
	fmt.Println(q + w + e + r)
	fmt.Println(t)

	// math包的常量
	fmt.Println(math.MaxFloat32)

	aa := 1
	bb := "2"
	// string()会把int转换成ASCII码，如果需要转换成 原值的string
	// 需要导入strconv包, Itoa代表int to string, Atoi相反
	fmt.Println(strconv.Itoa(aa) + bb)
}
