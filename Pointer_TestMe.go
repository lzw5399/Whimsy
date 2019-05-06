package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("mode", "sds", "process mode")

func main() {
	// flag.Parse()

	fmt.Println(mode)

	fmt.Println(*mode)
}
