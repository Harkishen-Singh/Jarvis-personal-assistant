package utils

import (
	"testing"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"fmt"
)

func Test_message_1_Service(t *testing.T) {

	data := map[string]string{"username": "default", "message": "Hi! this is test bot"}
	inJSON, _ := json.Marshal(data)
	// req, err := http.Post("http://127.0.0.1:3000/message", "application/x-www-form-urlencoded", bytes.NewBuffer(inJSON))
	req, err := http.NewRequest("POST", "http://127.0.0.1:3000/message", bytes.NewBuffer(inJSON))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	if !(string(body) == "{\"status\": \"success\"}") {
		t.Errorf("Response didnt match as required")
	}
}

func Test_message_2_Service(t *testing.T) {

	data := map[string]string{"username": "not_default", "message": "Hi! this is test bot part 2"}
	inJSON, _ := json.Marshal(data)
	// req, err := http.Post("http://127.0.0.1:3000/message", "application/x-www-form-urlencoded", bytes.NewBuffer(inJSON))
	req, err := http.NewRequest("POST", "http://127.0.0.1:3000/message", bytes.NewBuffer(inJSON))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	if !(string(body) == "{\"status\": \"success\"}") {
		t.Errorf("Response didnt match as required")
	} else {
		fmt.Println("Got response as => ", string(body))
	}
}
