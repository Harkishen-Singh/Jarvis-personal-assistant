package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/config"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/logger"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/messages"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/services/herokuhost"
	utils "github.com/Harkishen-Singh/Jarvis-personal-assistant/service/utils"
)

// AST controls the basic conversation flow
// of the general messages asked by the users to
// the personal assistant.
//
// This is filtered by various comparisions done
// with appropriate matchings and the required
// action is carried out.
func AST(routeObject response, w http.ResponseWriter) {

	var (
		matchPars, remainingString string
		a                          []string
	)

	message := routeObject.message
	low(&message)
	messageArr := strings.Fields(message)

	priority := config.Get().KeywordPriority

	for i := 0; i < len(messageArr); i++ {
		for _, prior := range priority {
			if messageArr[i] == prior {
				a = append(a, messageArr[i])
				if i < len(messageArr) {
					messageArr = append(messageArr[:i], messageArr[i+1:]...)
					i = i - 1
				} else {
					fmt.Println("Position Invalid")
				}
				break
			}
		}
	}

	if len(a) > 0 {
		sort := customSort(a, priority, len(a), len(priority))
		matchPars = sort[0]
		remainingString = strings.Join(messageArr[:], " ")
		messageArr = append([]string{matchPars}, messageArr...)
	} else {
		remainingString = strings.Join(messageArr[:], " ")
	}

	low(&matchPars)

	if utils.Connected() {

		if matchPars == "google" || matchPars == "search" {
			// Google search

			if len(remainingString) == 0 {
				remainingString = "google"
			}
			response, _ := processGoogleResponses(remainingString, "com", "en", nil, 1, 10, 5)
			responseJSON := jsonResponseQuery{
				Status:  true,
				Message: "here are the top search results",
				Result:  response,
			}

			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech(responseJSON.Message, 0)

		} else if matchPars == "yahoo" {
			// Yahoo search (planned to be deprecated in future releases)

			query := ""
			if len(remainingString) == 0 {
				query = "yahoo"
			} else {
				query = remainingString
			}

			result := scrapeYahoo(query)

			responseJSON := jsonResponseQuery{
				Status:  true,
				Message: "here are the top search results",
				Result:  result,
			}

			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech(responseJSON.Message, 0)

		} else if matchPars == "bing" {
			// Bing search (planned to be deprecated in future releases)

			if len(remainingString) == 0 {
				remainingString = "bing"
			}

			response, err := processBingResponses(remainingString, "com", nil, 1, 10, 5)
			if err != nil {
				logger.Error(err)
			}

			responseJSON := jsonResponseQuery{
				Status:  true,
				Message: "here are the top search results",
				Result:  response,
			}

			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech(responseJSON.Message, 0)

		} else if matchPars == "youtube" || matchPars == "videos" || matchPars == "watch" {
			// Requests for watching videos. Handled by Youtube search.

			query := ""
			if len(remainingString) == 0 {
				query = "https://www.youtube.com/results?search_query=youtube"
			} else {
				query = "https://www.youtube.com/results?search_query=" + remainingString
			}

			result := HandlerYoutube("GET", query)
			responseJSON := jsonResponseQuery{
				Status:  true,
				Message: "here are the top search videos",
				Result:  processYoutubeResponses(result),
			}

			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech(responseJSON.Message, 0)

		} else if matchPars == "images" || matchPars == "image" {
			// Images request. Handled by google images search.

			query := ""
			if len(remainingString) == 0 {
				query = "images"
			} else {
				query = remainingString
			}

			responseJSON := jsonResponseQuery{
				Status:  true,
				Message: "here are the searched images",
				Result:  scrapeImage(query),
			}

			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech(responseJSON.Message, 0)

		} else if matchPars == "weather" {

			if len(messageArr) < 3 {
				w.Write([]byte(`{"status": "success", "message": "ENTER: weather <city> <state>", "result": ""}`))
			} else {
				city, state := messageArr[len(messageArr)-2], messageArr[len(messageArr)-1]

				response := jsonResponseWeather{
					Status:  true,
					Message: "here are the current weather conditions",
					Result:  weather(city, state),
				}

				jData, _ := json.Marshal(response)
				w.Write(jData)
				TextToSpeech(response.Message+city+" "+state, 0)
			}
		} else if matchPars == "meaning" {

			if len(messageArr) == 1 {
				w.Write([]byte(`{"status": "success", "message": "ENTER: meaning <word>", "result": ""}`))
			} else {

				wordStr := remainingString

				// Confirm the below.
				if len(wordStr) > 0 {

					responseJSON := jsonResponseMeaning{
						Status:  true,
						Message: "here is the meaning of the searched word",
						Result:  meanings(wordStr),
					}

					jData, _ := json.Marshal(responseJSON)
					w.Write(jData)
					TextToSpeech(responseJSON.Message+" "+filterForSpeech(wordStr), 0)

				} else {
					response, err := processGoogleResponses(wordStr, "com", "en", nil, 1, 10, 5)
					if err != nil {
						logger.Error(err)
					}
					responseJSON := jsonResponseQuery{
						Status:  true,
						Message: "here are the top search results",
						Result:  response,
					}

					jData, _ := json.Marshal(responseJSON)
					w.Write(jData)
					TextToSpeech(responseJSON.Message+" "+filterForSpeech(wordStr), 0)
				}
			}
		} else if matchPars == "medicine" {
			if len(messageArr) <= 1 {
				w.Write([]byte(`{"status": "success", "message": "ENTER: medicine <generic / common name>", "result": ""}`))
			} else {
				med := messageArr[len(messageArr)-1]
				result := messages.HealthMedController(med, w)
				TextToSpeech(result, 0)
			}
		} else if matchPars == "symptoms" {
			// add support for multiple symptoms at once and use ML to determine the best medicine suited

			if len(messageArr) < 2 {
				w.Write([]byte(`{"status": "success", "message": "ENTER: symptoms <symptom / condition>", "result": ""}`))
			} else {
				fmt.Println("inside")
				symp := strings.Join(messageArr[1:len(messageArr)], " ")
				result := messages.HealthSympController(symp, w)
				TextToSpeech(result, 0)
			}
		} else if strings.HasPrefix(strings.ToLower(message), "set reminder") {

			w.Write([]byte(`{"status": "success", "message": "Enter Reminder details : ", "result": ""}`))

		} else if strings.HasPrefix(strings.ToLower(message), "show reminder") {

			responseJSON := reminderResponse{
				Status:  true,
				Message: "Here are your reminders : ",
				Result:  ShowReminder(),
			}

			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech("Here are your reminders.", 0)

		} else if strings.HasPrefix(strings.ToLower(message), "deploy") {
			// support for deployment functionality (deprecated)

			status := herokuhost.DeploymentFunction(messageArr[len(messageArr)-1], w)
			TextToSpeech(filterForSpeech(status), 0)

		} else if strings.HasPrefix(strings.ToLower(message), "send mail") {

			w.Write([]byte(`{"status": "success", "message": "Enter Mail details : ", "result": ""}`))

		} else {
			// general conversation
			speech := messages.GeneralConvHandler(routeObject.message, routeObject.username, w)
			TextToSpeech(filterForSpeech(speech), 0)
		}
	} else {

		low(&matchPars)
		low(&message)

		if matchPars == "google" || matchPars == "yahoo" || matchPars == "bing" || matchPars == "youtube" ||
			matchPars == "image" || matchPars == "weather" || matchPars == "medicine" || matchPars == "symptoms" ||
			strings.HasPrefix(message, "send mail") {

			w.Write([]byte(`{"status": "success", "message": "Services unavailable at the moment ! Check your Internet Connection and try again.", "result": ""}`))
			TextToSpeech("Services unavailable at the moment!", 0)

		} else if strings.HasPrefix(message, "set reminder") {

			w.Write([]byte(`{"status": "success", "message": "Enter Reminder details : ", "result": ""}`))

		} else if strings.HasPrefix(message, "show reminder") {

			responseJSON := reminderResponse{
				Status:  true,
				Message: "Here are your reminders : ",
				Result:  ShowReminder(),
			}

			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech("Here are your reminders.", 0)
		} else {
			// Conversation.

			speech := messages.GeneralConvHandler(routeObject.message, routeObject.username, w)
			TextToSpeech(filterForSpeech(speech), 0)
		}
	}
}

func low(s *string) {
	*s = strings.ToLower(*s)
}

// customSort() to sort an array according to the order defined by another array
func customSort(arr1 []string, arr2 []string, m, n int) []string {
	freq := make(map[string]int)

	for i := 0; i < m; i++ {
		freq[arr1[i]]++
	}

	index := 0

	for i := 0; i < n; i++ {
		for freq[arr2[i]] > 0 {
			arr1[index] = arr2[i]
			index++
			freq[arr2[i]]--
		}
		delete(freq, arr2[i])
	}
	return arr1
}
