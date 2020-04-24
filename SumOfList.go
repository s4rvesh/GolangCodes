package main

import (
	"fmt"
)

func sum(list []int) int {

	fmt.Println(list)
	if len(list) == 0 {

		return 0
	}
	return list[0] + sum(list[1:])

}

func main() {

	a := []int{1, 2, 3, 4, 50}

	i := sum(a)

	fmt.Printf("The sum is: %v", i)
}
