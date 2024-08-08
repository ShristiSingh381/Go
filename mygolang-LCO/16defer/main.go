package main

import "fmt"

func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("Hello")
	mydefer()
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

}
