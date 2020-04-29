package main

import (
	"fmt"
)

func main() {

	list := []int{12, 34, 566, 32, 33, 11, 55, 77}

	sorted := bubble(list)
	fmt.Println(sorted)
}

func bubble(list []int) []int {

	for sweep := 0; sweep < len(list); sweep++ {

		swapped := false

		for i := 0; i < len(list)-1-sweep; i++ {

			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}

	}

	return list

}
