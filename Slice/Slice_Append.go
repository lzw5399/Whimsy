package main

import "fmt"

func main() {
	/*s1 := make([]int, 3, 6)
	fmt.Printf("%p\n", s1) // 打印出内存地址
	s1 = append(s1, 1, 2, 3, 4)
	fmt.Printf("%p\n", s1) // 打印出内存地址*/

	a := []int{0, 1, 2, 3, 4}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(s1, s2)
	fmt.Printf("append之前的s2地址:%p,长度:%v,容量%v\n", s2, len(s2), cap(s2))

	s2 = append(s2, s1...)
	fmt.Printf("append之后的s2地址:%p,长度:%v,容量%v\n", s2, len(s2), cap(s2))
	s1[0] = 9
	fmt.Println(s1, s2)
	fmt.Println("此时原数组=", a)
}
