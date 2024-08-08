package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps learning")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "RUBY"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println("JS shorts for : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("For Key %v,Value is %v\n", key, value)
	}

}
