package main

import "fmt"

func main() {
	fmt.Println("Welcome to ifelse")

	logincount := 3
	var result string

	if logincount < 10 {
		result = "Regular User"
	} else if logincount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if num := 10; num < 10 {
		fmt.Println("No")
	} else {
		fmt.Printf("num is not less than %v\n", num)
	}
}
