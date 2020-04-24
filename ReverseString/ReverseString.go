package main

import (
	"fmt"
)

func rev(str string) string {

	var res string

	for i := 0; i < len(str); i++ {

		res = string(str[i]) + res
	}

	return res
}

func main() {

	s := "sarvesh"

	r := rev(s)
	fmt.Printf("hello: %v", r)

}
