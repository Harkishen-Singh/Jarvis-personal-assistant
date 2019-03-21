package controllers

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type response struct {
	username string
	message string
}

type statusCode struct {
	status string `json:"status"`
}

// MessagesController controls messages handling
func MessagesController(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()
	contentReceived := response{
		username: r.FormValue("username"),
		message: r.FormValue("message"),
	}
	fmt.Println(contentReceived)
	reponseToReq := &statusCode{status: "success"}
	resp, _ := json.Marshal(reponseToReq)
	fmt.Println(resp)
	w.Write(resp)
}