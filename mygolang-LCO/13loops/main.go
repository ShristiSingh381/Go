package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops learning")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for i, day := range days {
	// 	fmt.Printf("Index is %v and value is %v\n", i, day)
	// }

	i := 1

	for i < 10 {
		fmt.Println(i)
		i++
	}
}
