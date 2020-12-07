package utils

import (
	"net/http"
)

// Connected checks whether the application is connected to the internet
// returns the status as a boolean value
func Connected() bool {
	_, err := http.Get("http://www.google.co.in")
	return err == nil
}
