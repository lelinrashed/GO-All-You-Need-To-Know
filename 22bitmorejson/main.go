package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")

	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{
			"React js Bootcamp",
			299,
			"learncodeonline.in",
			"123456",
			[]string{"web-dev", "react", "bootcamp"},
		},
		{
			"Vue js Bootcamp",
			299,
			"learncodeonline.in",
			"123456",
			[]string{"web-dev", "vue", "bootcamp"},
		},
		{
			"Angular js Bootcamp",
			299,
			"learncodeonline.in",
			"123456",
			nil,
		},
	}

	// package this data as json data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
                "coursename": "Vue js Bootcamp",
                "Price": 299,
                "website": "learncodeonline.in",
                "tags": [
                        "web-dev",
                        "vue",
                        "bootcamp"
                ]
        }	
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)

		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	// some cases where we need to add data as key value pair
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is %T\n", k, v, v)
	}
}
