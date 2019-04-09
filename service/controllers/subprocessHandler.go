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

// HandlerGoogle handles the subprocesses related to fetchdata_query.js for google search operations
// returns the http body as string
func HandlerGoogle(method string, url string) string {

	directory, _ := os.Getwd()
	fmt.Println("google-query request")
	fmt.Println("method -> " + method + " url -> " + url + " direc -> " + directory)
	result, err := exec.Command("node", "subprocesses/fetchdata_query.js", method, url).Output()
	// result, err := exec.Command("pwd").Output()
	if err != nil {
		panic(err)
	}
	return string(result)

}

// HandlerYahoo handles the subprocesses related to fetchdata_query.js for yahoo search operations
// returns the http body as string
func HandlerYahoo(method string, url string) string {

	directory, _ := os.Getwd()
	fmt.Println("yahoo-query request")
	fmt.Println("method -> " + method + " url -> " + url + " direc -> " + directory)
	result, err := exec.Command("node", "subprocesses/fetchdata_query.js", method, url).Output()
	// result, err := exec.Command("pwd").Output()
	if err != nil {
		panic(err)
	}
	return string(result)

}

// HandlerBing handles the subprocesses related to fetchdata_query.js for bing search operations
// returns the http body as string
func HandlerBing(method string, url string) string {

	directory, _ := os.Getwd()
	fmt.Println("bing-query request")
	fmt.Println("method -> " + method + " url -> " + url + " direc -> " + directory)
	result, err := exec.Command("node", "subprocesses/fetchdata_query.js", method, url).Output()
	// result, err := exec.Command("pwd").Output()
	if err != nil {
		panic(err)
	}
	return string(result)

}

// HandlerWeather handles the subprocesses related to weather.py for weather operations
// returns the http body as string
func HandlerWeather(city string, state string) string {

	fmt.Println("weather request")
	fmt.Println("city -> ", city, " state -> ", state)
	result, err := exec.Command("python", "subprocesses/weather.py", city, state).Output()
	if err != nil {
		fmt.Println("Seems like python version 2 is not yet installed or the pip dependencies are not installed")
		panic(err)
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
		panic(err)
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
		panic(err)
	}
	return string(result)

}

// HandlerMeaning handles the subprocesses related to dictionary.py for meaning operations
// returns the http body as string
func HandlerMeaning(word string) string {

	fmt.Println("meaning request")
	fmt.Println("word -> ", word)
	result, err := exec.Command("python", "subprocesses/dictionary.py", word).Output()
	if err != nil {
		fmt.Println("Seems like python version 2 is not yet installed or the pip dependencies are not installed")
		panic(err)
	}
	return string(result)

}

