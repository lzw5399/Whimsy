package main

import (
	"fmt"
)

func main() {
	a.ChangeNumber()
	fmt.Println(a)
}

type A struct {
	Number int
}

func (a A) ChangeNumber() {
	a.Number = 66
}
func (a *A) ChangeNumber() {

}
