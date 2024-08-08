package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&payment=ncjlkfnvwj"

func main() {
	fmt.Println("Welcome to urls in Golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of qparams is %T\n", qparams)

	fmt.Println(qparams["payment"])

	for _, abc := range qparams {
		fmt.Println("Param is ", abc)
	}

}
