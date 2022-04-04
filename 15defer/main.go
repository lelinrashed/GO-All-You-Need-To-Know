package main

import "fmt"

func main() {
	fmt.Println("Defer in Golang")

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hello")
}
