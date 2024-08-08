package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")
	shristi := User{"Shristi", "shristi@go.dev", true, 23}
	fmt.Printf("Value of Shristi is %+v\n", shristi)
	fmt.Println(shristi.Name)
	shristi.GetStatus()
	fmt.Println(shristi.EmailID)
	shristi.NewMail()

}

type User struct {
	Name    string
	EmailID string
	Status  bool
	Age     int
}

func (u User) GetStatus() {
	fmt.Println("Is user active ", u.Status)
}

func (u User) NewMail() {
	u.EmailID = "shri@go.dev"
	fmt.Println("User email is : ", u.EmailID)

}
