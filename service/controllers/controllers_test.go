package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestGoogleSearch(t *testing.T) {
	os.Setenv("ENV", "test")
	form := url.Values{}
	form.Add("username", "default")
	form.Add("message", "google CET bhubaneswar")
	resp, err := http.PostForm("http://localhost:3000/message", form)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	// success check
	if !(strings.Contains(string(body), "{\"status\":\"success\"") || strings.Contains(string(body), "{\"status\":true")) {
		t.Errorf("Response didnt match as required")
	}

	// incoming data checks
	if !strings.Contains(string(body), "here are the top search results") {
		t.Errorf("Google functionality response didn't match as required!")
	}

}

func disableTestYahooSearch(t *testing.T) {

	form := url.Values{}
	form.Add("username", "default")
	form.Add("message", "yahoo CET bhubaneswar")
	resp, err := http.PostForm("http://localhost:3000/message", form)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	// success check
	if !(strings.Contains(string(body), "{\"status\":\"success\"") || strings.Contains(string(body), "{\"status\":true")) {
		t.Errorf("Yahoo Response didnt match as required")
	}

	// incoming data checks
	if !strings.Contains(string(body), "here are the top search results") {
		t.Errorf("Yahoo functionality response didn't match as required!")
	}

}

func TestBingSearch(t *testing.T) {
	os.Setenv("ENV", "test")
	form := url.Values{}
	form.Add("username", "default")
	form.Add("message", "bing CET bhubaneswar")
	resp, err := http.PostForm("http://localhost:3000/message", form)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	// success check
	if !(strings.Contains(string(body), "{\"status\":\"success\"") || strings.Contains(string(body), "{\"status\":true")) {
		t.Errorf("Bing Response didnt match as required")
	}

	// incoming data checks
	if !strings.Contains(string(body), "here are the top search results") {
		t.Errorf("Bing functionality response didn't match as required!")
	}

}

// TODO

func disableTestImagesSearch(t *testing.T) {

	form := url.Values{}
	form.Add("username", "default")
	form.Add("message", "image CET bhubaneswar")
	resp, err := http.PostForm("http://localhost:3000/message", form)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	// success check
	if !(strings.Contains(string(body), "{\"status\":\"success\"") || strings.Contains(string(body), "{\"status\":true")) {
		t.Errorf("Google Image Response didnt match as required")
	}

	// incoming data checks
	if !strings.Contains(string(body), "here are the searched images") {
		t.Errorf("Google image functionality response didn't match as required!")
	}

}

func disableTestYoutube(t *testing.T) {

	form := url.Values{}
	form.Add("username", "default")
	form.Add("message", "youtube palazzo")
	resp, err := http.PostForm("http://localhost:3000/message", form)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	// success check
	if !(strings.Contains(string(body), "{\"status\":\"success\"") || strings.Contains(string(body), "{\"status\":true")) {
		t.Errorf("Youtube Response didnt match as required")
	}

	// incoming data checks
	if !strings.Contains(string(body), "here are the top search videos") {
		t.Errorf("Youtube functionality response didn't match as required!")
	}

}

func disableTestWeather_1(t *testing.T) {

	form := url.Values{}
	form.Add("username", "default")
	form.Add("message", "weather bhubaneswar odisha")
	resp, err := http.PostForm("http://localhost:3000/message", form)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	// success check
	if !(strings.Contains(string(body), "{\"status\":\"success\"") || strings.Contains(string(body), "{\"status\":true")) {
		t.Errorf("Weather Response 1 didnt match as required")
	}

	// incoming data checks
	if !strings.Contains(string(body), "here are the current weather conditions") {
		t.Errorf("Weather functionality 1 response didn't match as required!")
	}

}

// func TestWeather_2(t *testing.T) {

// 	form := url.Values{}
// 	form.Add("username", "default")
// 	form.Add("message", "weather in bhubaneswar odisha")
// 	resp, err := http.PostForm("http://localhost:3000/message", form)
// 	if err != nil {
// 		panic(err)
// 	}
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Println(string(body))
// 	// success check
// 	if !(strings.Contains(string(body), "{\"status\":\"success\"") || strings.Contains(string(body), "{\"status\":true")) {
// 		t.Errorf("Weather Response 2 didnt match as required")
// 	}

// 	// incoming data checks
// 	if !strings.Contains(string(body), "here are the current weather conditions") {
// 		t.Errorf("Weather functionality 2 response didn't match as required!")
// 	}

// }
