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

// HandlerGoogle handles the subprocesses related to fetchdata_query.js
// returns the http body as string
func HandlerGoogle(method string, url string) string {

	directory, _ := os.Getwd()
	fmt.Println("method -> " + method + " url -> " + url + " direc -> " + directory)
	result, err := exec.Command("node", "subprocesses/fetchdata_query.js", method, url).Output()
	// result, err := exec.Command("pwd").Output()
	if err != nil {
		panic(err)
	}
	return string(result)

}

// HandlerYahoo handles the subprocesses related to fetchdata_query.js
// returns the http body as string
func HandlerYahoo(method string, url string) string {

	directory, _ := os.Getwd()
	fmt.Println("method -> " + method + " url -> " + url + " direc -> " + directory)
	result, err := exec.Command("node", "subprocesses/fetchdata_query.js", method, url).Output()
	// result, err := exec.Command("pwd").Output()
	if err != nil {
		panic(err)
	}
	return string(result)

}

// HandlerBing handles the subprocesses related to fetchdata_query.js
// returns the http body as string
func HandlerBing(method string, url string) string {

	directory, _ := os.Getwd()
	fmt.Println("method -> " + method + " url -> " + url + " direc -> " + directory)
	result, err := exec.Command("node", "subprocesses/fetchdata_query.js", method, url).Output()
	// result, err := exec.Command("pwd").Output()
	if err != nil {
		panic(err)
	}
	return string(result)

}