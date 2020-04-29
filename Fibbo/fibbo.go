package main

import (
	"fmt"
)

/*func fibbo(n int) int {

	if n <= 1 {

		return n

	}

	return fibbo(n-1) + fibbo(n-2)
}*/

func fibbo(N int) int {
	if N <= 1 {

		return N

	}

	nMin2 := 0
	nMin1 := 1
	var cur int

	for i := 2; i <= N; i++ {

		cur = nMin2 + nMin1
		nMin2 = nMin1
		nMin1 = cur
	}

	return cur

}

func main() {

	for i := 0; i < 20; i++ {
		a := fibbo(i)
		fmt.Println(a)
	}

}
