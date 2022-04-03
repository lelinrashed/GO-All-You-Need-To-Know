package main

import "fmt"

func main() {
	fmt.Println("Welcome to slice in golang")

	var fruitList = []string{}

	fruitList = append(fruitList, "Apple", "Banana", "Tomato", "Peach")

	fruitList = append(fruitList[1:])

	fmt.Println(fruitList)

	highScores := make([]string, 0)

	highScores = append(highScores, "banana")

	fmt.Println(highScores)
}
