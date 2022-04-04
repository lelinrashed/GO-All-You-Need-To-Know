package main

import "fmt"

func main() {
	fmt.Println("Method in golang")

	rashed := User{"rashed", "rashed@gmail.com", true, 26}

	// fmt.Println(rashed)
	// fmt.Printf("%+v\n", rashed)

	rashed.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active", u.Status)
}
