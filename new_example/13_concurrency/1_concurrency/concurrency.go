package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// go Go()
	go WriteWithFileWrite()
	time.Sleep(2 * time.Second)
}

func Go() {
	fmt.Println("GO GO GO!!")
}

func WriteWithFileWrite() {
	name := "D:\\GoCode\\src\\GolangDemo\\new_example\\13_concurrency\\1_concurrencytestwritefile.txt"
	content := "Hello, xxbandy.github.io!"
	fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	defer fileObj.Close()
	if _, err := fileObj.WriteString(content); err == nil {
		fmt.Println("Successful writing to the file with os.OpenFile and *File.WriteString method.", content)
	}
	contents := []byte(content)
	if _, err := fileObj.Write(contents); err == nil {
		fmt.Println("Successful writing to thr file with os.OpenFile and *File.Write method.", content)
	}
}
