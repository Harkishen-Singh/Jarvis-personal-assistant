package controllers

import (
	"net/http"
	"../utils"
	"fmt"
	"encoding/json"
)

type response struct {
	username string
	message string
}

// MessagesController controls messages handling
func MessagesController(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	utils.LoggerWarn("message from user")
	body := json.NewDecoder(r.Body)
	var resp response
	body.Decode(&resp)
	fmt.Println(body)
	fmt.Println(resp)

}