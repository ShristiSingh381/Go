package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println(result)

	pro := add(2, 6, 8, 4)
	fmt.Println(pro)

}

func add(val ...int) int {
	total := 0
	for _, val := range val {
		total += val
	}
	return total
}

func adder(a int, b int) int {
	return a + b
}

func greeter() {
	fmt.Println("Namaste from Golang")
}
