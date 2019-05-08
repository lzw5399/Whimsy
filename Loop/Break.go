package Loop

import "fmt"

func main() {
label1:
	for {
		for a := 1; a < 10; a++ {
			if a > 3 {
				break label1
			}
			fmt.Println(a)
		}
	}
}
