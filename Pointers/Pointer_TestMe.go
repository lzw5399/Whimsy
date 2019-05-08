package Pointers

import "fmt"

// var mode = flag.String("mode", "sds", "process mode")

func main() {
	a := new(int)
	*a = 88
	fmt.Println(*a)
}
