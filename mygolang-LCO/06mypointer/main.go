package main

import "fmt"

func main() {
	fmt.Println("Welcome to learn Pointers")

	// var ptr *int
	// fmt.Println("Value of ptr is : ", ptr)

	mynumber := 23

	var ptr = &mynumber
	fmt.Println("Value of ptr is", *ptr)
	fmt.Println("Address of ptr is", ptr)

	*ptr = *ptr + 2
	fmt.Println("Value of mynumber is", mynumber)

}
