/**
* This handler controls operations of selenium core,
* which could be used to scrap any website till date,
* if designed sensibily.
 */

package controllers

import (
	"fmt"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/logger"
	"os"
	"os/exec"
)

// HandlerYahoo handles the subprocesses related to fetchdata_query.js for yahoo search operations
// returns the http body as string
func HandlerYahoo(method string, url string) string {

	directory, _ := os.Getwd()
	fmt.Println("yahoo-query request")
	fmt.Println("method -> " + method + " url -> " + url + " direc -> " + directory)
	result, err := exec.Command("node", "subprocesses/fetchdata_query.js", method, url).Output()
	// result, err := exec.Command("pwd").Output()
	if err != nil {
		logger.Error(err)
	}
	return string(result)

}

// HandlerYoutube handles the subprocesses related to fetchdata_query.js for youtube operations
// returns the http body as string
func HandlerYoutube(method string, url string) string {

	directory, _ := os.Getwd()
	fmt.Println("youtube-query request")
	fmt.Println("method -> " + method + " url -> " + url + " direc -> " + directory)
	result, err := exec.Command("node", "subprocesses/fetchdata_query.js", method, url).Output()
	// result, err := exec.Command("pwd").Output()
	if err != nil {
		logger.Error(err)
	}
	return string(result)

}

// HandlerImage handles the subprocesses related to fetchdata_query.js for image operations
// returns the http body as string
func HandlerImage(method string, url string) string {

	directory, _ := os.Getwd()
	fmt.Println("google-image-query request")
	fmt.Println("method -> " + method + " url -> " + url + " direc -> " + directory)
	result, err := exec.Command("node", "subprocesses/fetchdata_query.js", method, url).Output()
	if err != nil {
		logger.Error(err)
	}
	return string(result)

}
