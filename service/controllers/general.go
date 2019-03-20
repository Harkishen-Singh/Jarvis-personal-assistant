package controllers

import (
	"net/http"
	"fmt"
)

// HomeController controls homepage execution
func HomeController(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hey user! -- from Jarvis %s", r.URL.Path[1:])
}