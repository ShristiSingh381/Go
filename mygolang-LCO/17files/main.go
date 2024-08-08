package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in go")

	content := "This is needs to go in a file - LearnGo.in"
	file, err := os.Create("./mylcogofilelist.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	fmt.Println(length)
	defer file.Close()
	readfile("/home/shristi/Downloads/golang-LCO/mygolang-LCO/17files/mylcogofilelist.txt")
}

func readfile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	checkNilError(err)

	fmt.Println("Text present inside the file \n", string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
