package main

import "fmt"

func main() {
	fmt.Println("Welcome to study arrays")

	var fruitlist [4]string

	fruitlist[0] = "Banana"
	fruitlist[1] = "Apple"
	fruitlist[3] = "Guava"

	fmt.Println("My fruitlist contains : ", fruitlist)
	fmt.Println("length of myfruitlist is : ", len(fruitlist))

	var vegList = [3]string{"lauki", "brinjal", "ladyfinger"}
	fmt.Println("My veglist contains : ", vegList)
	fmt.Println("length of veglist is : ", len(vegList))

}
