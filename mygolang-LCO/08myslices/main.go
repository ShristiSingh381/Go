package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices learning")

	var fruitlist = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Banana")
	fmt.Println(fruitlist)

	// fruitlist = append(fruitlist[1:4])
	// fmt.Println(fruitlist)

	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 235
	highscores[2] = 236
	highscores[3] = 237
	// highscores[4] = 238

	highscores = append(highscores, 55, 66, 99)

	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	// var courses = []string{"Go", "jenkins", "AWS", "EKS", "CICD"}
	// fmt.Println(courses)
	// var index int = 2
	// courses = append(courses[:index], courses[index+1:]...)
	// fmt.Println(courses)

}
