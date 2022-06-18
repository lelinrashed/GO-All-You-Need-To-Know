package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lelinrashed/mongoapi/router"
)

func main() {
	fmt.Println("Mongodb API")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")
}
