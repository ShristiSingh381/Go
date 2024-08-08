package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.w3schools.com/"

func main() {
	fmt.Println("Welcome to handling web request in Golang")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", response)

	defer response.Body.Close()
	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databyte)
	fmt.Println(content)
}
