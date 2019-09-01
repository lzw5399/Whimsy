package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8880/user/login")
	handleError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)
	fmt.Println(string(body))
}

func handleError(err Error) {
	if err != nil {
		fmt.Println("failed: ", err)
	}
}

type Error interface {
	Error() string
}
