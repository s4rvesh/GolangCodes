package main

import "fmt"

func numl(list []int, num int) bool {

	for _, i := range list {

		if i == num {

			return true

		}
	}
	return false
}

func main() {

	a := []int{10, 9, 6, 7, 8}
	b := numl(a, 4)

	if b == true {
		fmt.Println("number there")
	} else {
		fmt.Println("Not there")
	}

}
