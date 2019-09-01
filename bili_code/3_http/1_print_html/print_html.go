package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://go.codepie.fun")
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
