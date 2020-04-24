package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(num int) {

	for i := 1; i <= num; i++ {

		var output string = ""

		if i%3 == 0 {
			output += string("Fizz")
		}
		if i%5 == 0 {
			output += string("Buzz")

		}
		if output == "" {
			output += strconv.Itoa(i)
		}

		fmt.Printf(output)
		fmt.Printf("  ")

	}
}

func main() {
	a := 15
	fizzbuzz(a)
}
