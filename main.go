package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!!")
	resp, err := http.Get("http://www.example.com")
	if err != nil {
		panic("I do not understand Japanese.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("I do not understand English.")
	}
	fmt.Printf("body = %T\n", body)
	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Println(string(body))
}
