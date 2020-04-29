package main

import (
	"fmt"
)

func dectobase(dec, base int) string {

	var res string = ""

	for dec > 0 {

		const charset = "012346789ABCDEF"

		a := dec % base
		dec = dec / base

		res = string(charset[a]) + res

	}

	return res
}

func main() {

	base := 16
	dec := 15

	ans := dectobase(dec, base)

	fmt.Printf("value: %v", ans)
}
