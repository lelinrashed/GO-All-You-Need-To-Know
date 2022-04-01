package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input : !"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for out pizza: ")

	// comma ok syntaxt
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
}
