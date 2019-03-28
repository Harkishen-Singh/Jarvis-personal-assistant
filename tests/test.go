package main

import (
	"fmt"
	"os/exec"
)


func main() {

	out, err := exec.Command("mocha").Output()
	if err != nil {
		panic (err)
	}
	fmt.Println(string(out))
	fmt.Println(len(out))
}