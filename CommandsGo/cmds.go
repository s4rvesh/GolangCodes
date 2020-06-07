package main

import (
	"fmt"
	//"runtime"
	"os/exec"
)

func execute() {

	out, err := exec.Command("ls", "-ltr").Output()

	if err != nil {

		fmt.Printf("%s", err)
	}

	fmt.Println("Command Successfully Executed")

	output := string(out[:])

	fmt.Printf("%s", output)
}

func main() {

	/*if runtime.GOOS == "windows" {

		fmt.Printf("no support")
	} else {*/
	execute()
	//}
}
