package main

import (
	"fmt"
)

func fact(primes []int, num int) []int {

	var res []int

	for _, prim := range primes {

		for num%prim == 0 {

			res = append(res, prim)
			num = num / prim

		}
	}

	if num > 1 {

		res = append(res, num)
	}

	return res

}

func main() {

	prim := []int{2, 3, 4, 5, 7, 11}

	res := fact(prim, 28)

	fmt.Println(res)
}
