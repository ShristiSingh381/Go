package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")

	shristi := User{"Shristi", "shristi@go.dev", true, 23}
	fmt.Printf("Value of Shristi is %+v\n", shristi)
	fmt.Printf(shristi.Name)
}

type User struct {
	Name    string
	EmailID string
	Member  bool
	Age     int
}
