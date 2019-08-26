package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	ft := t.Format(time.ANSIC)
	fmt.Println(ft)
}
