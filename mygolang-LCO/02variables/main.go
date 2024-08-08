package main

import "fmt"

func main() {
	var username string = "Shristi"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var number int8 = 78
	fmt.Println(number)
	fmt.Printf("Variable is of type: %T \n", number)

	var smallFloat float32 = 256.35257628732897
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	numberofUser := 23455
	fmt.Println(numberofUser)
}
