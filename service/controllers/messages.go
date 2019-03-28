package controllers

import (
	"net/http"
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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()
	w.Write([]byte(`{"status": "success", "message": "Hi from reply bot"}`))

}