package main

import (
	"encoding/json"
	"fmt"
)

type pook struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name      string `json:"string"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {

	fmt.Println("Hello WOrld!!")

	auth := Author{Name: "Nawaz", Age: 24, Developer: true}

	book := pook{Title: "WAsseyPur", Author: auth}

	//fmt.Printf("%+v\n", book)

	byteArray, err := json.MarshalIndent(book, "", "  ")

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(string(byteArray))

}
