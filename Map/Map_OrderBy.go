package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{1: "a", 5: "e", 3: "c", 2: "b", 4: "d"}
	s := make([]int, len(m))
	i := 0
	for k := range m {
		s[i] = k
		i++
	}
	// 成功对key进行排序，可见出slice是引用类型，因为没有返回值
	sort.Ints(s)

	for j := 0; j < 5; j++ {
		fmt.Println(m[s[j]])
	}
	fmt.Println(s)
}
