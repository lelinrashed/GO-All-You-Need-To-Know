package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Struct in golang")

	rashed := User{"rashed", "rashed@gmail.com", true, 26}

	fmt.Println(rashed)
	fmt.Printf("%+v\n", rashed)
}
