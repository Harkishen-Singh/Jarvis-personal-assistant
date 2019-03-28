package controllers

import (
	"net/http"
	"strings"
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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()

	request := response{
		username: r.FormValue("username"),
		message: r.FormValue("message"),
	}
	fmt.Println(request)

	routes(request, w)

	w.Write([]byte(`{"status": "success", "message": "Hi from reply bot"}`))

}

func routes(routeObject response, w http.ResponseWriter) {

	message := routeObject.message
	fmt.Println(message)
	// messageTemp := message
	var firstPars string
	if strings.Contains(message, " ") {
		firstPars = message[:strings.Index(message, " ")]
	} else {
		firstPars = message
	}

	strArr := strings.Split(firstPars, " ")
	strArrDiff := strings.Split(message, " ")

	messageExceptFirstPars := strings.Join(stringDifference(strArr, strArrDiff), " ")
	// lastParsArr := strings.Split(messageTemp, " ")
	// lastPars := lastParsArr[len(lastParsArr) - 1]

	// single word operations

	if firstPars == "google" { // for google search

		query := "https://www.google.co.in/search?q=" + messageExceptFirstPars
		result := HandlerGoogle("GET", query)
		fmt.Println(result)

	}

}

// gives the difference of two string arrays as an array of the differed element
func stringDifference(slice1 []string, slice2 []string) []string {
    var diff []string

    // Loop two times, first to find slice1 strings not in slice2,
    // second loop to find slice2 strings not in slice1
    for i := 0; i < 2; i++ {
        for _, s1 := range slice1 {
            found := false
            for _, s2 := range slice2 {
                if s1 == s2 {
                    found = true
                    break
                }
            }
            // String not found. We add it to return slice
            if !found {
                diff = append(diff, s1)
            }
        }
        // Swap the slices, only if it was the first loop
        if i == 0 {
            slice1, slice2 = slice2, slice1
        }
    }

    return diff
}