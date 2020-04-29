package main

import "fmt"

func findint(a []int, sum int) (int, int) {

	for i, val := range a {

		for j, val2 := range a {

			if i == j {
				continue
			}
			if val+val2 == sum {

				return i, j
			}
		}
	}

	return -1, -1
}

func main() {

	list := []int{1, 2, 3, 4, 5, 6}
	sum := 7
	a, b := findint(list, sum)
	fmt.Printf("the number position: %v %v", a, b)
}
