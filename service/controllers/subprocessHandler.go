/**
* This handler controls operations of selenium core,
* which could be used to scrap any website till date,
* if designed sensibily.
*/

package controllers

import (
	"os/exec"
	"fmt"
	"os"
)

// HandlerGoogle handles the subprocesses related to google.js
// returns the http body as string
func HandlerGoogle(method string, url string) string {

	directory, _ := os.Getwd()
	fmt.Println("method -> " + method + " url -> " + url + " direc -> " + directory)
	result, err := exec.Command(directory + "/subprocesses", "node", "google.js", method, url).Output()
	// result, err := exec.Command("pwd").Output()
	if err != nil {
		panic(err)
	}
	return string(result)

}