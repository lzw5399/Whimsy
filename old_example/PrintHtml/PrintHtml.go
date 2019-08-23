package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	r, err := http.Get("https://github.com")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("body = %s\n", string(body))
	defer r.Body.Close()
}
