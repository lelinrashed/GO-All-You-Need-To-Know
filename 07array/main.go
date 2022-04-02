package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [5]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Mango"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))
}
