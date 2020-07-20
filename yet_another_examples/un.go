package main

import (
	"encoding/json"
	"fmt"
)

type ToolStruct struct {
	Count int `json:"count"`
	Page  int `json:"page"`
}

func main() {
	str := "{\"page\":1,\"count\":2}"
	var obj ToolStruct
	json.Unmarshal([]byte(str), &obj)

	fmt.Println(obj)
}

func methodd(p1 string, p2 int) (out1 bool, out2 rune) {
	return
}

func (t ToolStruct) extensionss(p1 string, p2 int) (out1 bool, out2 rune) {
	t.Count = p2

	out1 = true
	out2 = 12

	return
}
