package controllers

import (
	"net/http"
	"fmt"
)

type response struct {
	username string
	message string
}

type statusCode struct {
	status string
}

// MessagesController controls messages handling
func MessagesController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("thisss")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()
	w.Write([]byte(`{"status": "succe222ss", "message": "Hi from reply bot"}`))

}