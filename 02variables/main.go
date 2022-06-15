package main

import "fmt"

func main() {

	// Variable types in golang
	// int
	// float32
	// string
	// bool

	// In golang we can define variable through two different way
	// 1. with var keyword
	// var variableName type = value
	// 2. Another way to use ":=" this sign to define variable
	// variableName := value ( In this case type will be inferred by the value )

	// var username string = "rashed"
	// rashed := "lelin"
	// fmt.Printf(username)
	// fmt.Println(rashed)

	var (
		a int
		b int    = 1
		c string = "hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
