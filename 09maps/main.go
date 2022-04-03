package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["go"] = "Golang"

	fmt.Println("js shorts for", languages["js"])

	// delete(languages, "js")

	fmt.Println("List of all langalanguages", languages)

	// Loops in map
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
