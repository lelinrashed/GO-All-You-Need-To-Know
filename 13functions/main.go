package main

import "fmt"

func main() {
	fmt.Println("Function in golang")
	greeter()

	result, message := proAdder(3, 4, 5, 7)
	fmt.Println(result, message)
}

func greeter() {
	fmt.Println("Hello")
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "this is total"
}
