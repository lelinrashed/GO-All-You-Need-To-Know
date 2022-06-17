package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj3435"

func main() {
	fmt.Println("Welcome to handling URLs to golang")
	fmt.Println(myUrl)

	//
	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)

}
