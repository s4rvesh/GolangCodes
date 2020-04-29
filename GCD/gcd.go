package main

import (
	"fmt"
)

func main() {

	a := 15
	b := 45

	n := gcd(a, b)

	fmt.Println(n)
}

func gcd(a, b int) int {

	for b != 0 {

		a, b = b, a%b
	}
	return a
}
