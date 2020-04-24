package main

import "fmt"

func fizzbuzz(num int) {

	for i := 1; i <= num; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Printf("Buzz")
		} else if i%3 == 0 {
			fmt.Printf("Fizz")
		} else {
			fmt.Printf("%v", i)
		}

		fmt.Printf("  \n")

	}
}

func main() {
	a := 15
	fizzbuzz(a)
}
