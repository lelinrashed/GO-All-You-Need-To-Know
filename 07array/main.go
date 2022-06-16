package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	// var fruitList [5]string

	// fruitList[0] = "Apple"
	// fruitList[1] = "Banana"
	// fruitList[3] = "Mango"

	// fmt.Println("Fruit list is: ", fruitList)
	// fmt.Println("Fruit list is: ", len(fruitList))

	// Fixed length
	// var arr1 = [3]int{1, 2, 3}
	// arr2 := [3]int{1, 3, 4}

	// fmt.Println(arr1)
	// fmt.Println(arr2)

	// Inferred length
	var arr1 = [...]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}
