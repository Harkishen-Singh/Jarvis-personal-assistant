package controllers

import (
	"net/http"
	"strings"
	"encoding/json"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/messages"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/services/herokuhost"
	"fmt"
)

type response struct {
	username string
	message string
}

type statusCode struct {
	status string
}

type messageQueryBody struct {
	Head string `json:"head"`
	Link string `json:"link"`
	Desc string `json:"desc"`
	DescLink string `json:"dlink"`
}

// type imageQueryBody struct{
// 	Link string `json:"link"`
// }

type reminderResponse struct {
	Status bool	`json:"status"`
	Message string `json:"message"`
	Result []reminder `json:"result"`
}

type jsonResponseQuery struct {
	Status bool	`json:"status"`
	Message string `json:"message"`
	Result []messageQueryBody `json:"result"`
}

// type jsonResponseImage struct {
// 	Status bool	`json:"status"`
// 	Message string `json:"message"`
// 	Result []imageQueryBody `json:"result"`
// }

type jsonResponseWeather struct {
	Status bool	`json:"status"`
	Message string `json:"message"`
	Result string `json:"result"`
}

type weatherStr struct {
	Time string `json:"time"`
	City string `json:"city"`
	Temperature string `json:"temperature"`
	DewPoint string `json:"dew_point"`
	Humidity string `json:"humidity"`
	Visibility string `json:"visibility"`
	FeelsLike string `json:"feels_like"`
}

type meaningStr struct {
	Meaning string `json:"meaning"`
	Example string `json:"example"`
	Submeaning []submeanStr `json:"submeaning"`
}

type submeanStr struct {
	Smean string
	Subexample string
}

type jsonResponseMeaning struct {
	Status bool	`json:"status"`
	Message string `json:"message"`
	Result []meaningStr `json:"result"`
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

}

func routes(routeObject response, w http.ResponseWriter) {

	message := routeObject.message
	message = strings.ToLower(message)
	messageArr := strings.Fields(message)
	var a []string
	priority := []string{"images", "image", "video", "videos", "watch", "youtube", "symptoms", "medicine", "weather", 
						 "meaning", "google", "yahoo", "bing", "search"}
	for i := 0; i < len(messageArr); i++ {
		for _, prior := range priority {
			if (messageArr[i] == prior) {
				a = append(a, messageArr[i])
				if i < len(messageArr) {
					messageArr = append(messageArr[:i], messageArr[i+1:]...)
					i = i - 1
				} else {
					fmt.Println("Position Invalid")
				}
				break;
			}
		}
	}
	matchPars := ""
	remainingString := ""
	if len(a) > 0 {
		sort := customSort(a, priority, len(a), len(priority))
		matchPars = sort[0]
		remainingString = strings.Join(messageArr[:]," ")
		messageArr = append([]string{matchPars}, messageArr...)
	} else {
		remainingString = strings.Join(messageArr[:], " ")
	}


	// single word operations


	if Connected() {

		if strings.ToLower(matchPars) == "google" || strings.ToLower(matchPars) == "search" { // for google search
			query := ""
			if len(remainingString) == 0 {
				query = "https://www.google.co.in/search?q=google"
			} else {
				query = "https://www.google.co.in/search?q=" + remainingString
			}
			result := HandlerGoogle("GET", query)

			// processing

			response := processGoogleResponses(result)
			responseJSON := jsonResponseQuery {
				Status: true,
				Message: "here are the top search results",
				Result: response,
			}
			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech(responseJSON.Message, 0)

		} else if strings.ToLower(matchPars) == "yahoo" {
			query := ""
			if len(remainingString) == 0 {
				query = "https://in.search.yahoo.com/search?p=yahoo"
			} else {
				query = "https://in.search.yahoo.com/search?p=" + remainingString
			}

			result := HandlerYahoo("GET", query)

			// processing

			response := processYahooResponses(result)
			responseJSON := jsonResponseQuery {
				Status: true,
				Message: "here are the top search results",
				Result: response,
			}
			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech(responseJSON.Message, 0)

		} else if strings.ToLower(matchPars) == "bing" {
			query := ""
			if len(remainingString) == 0 {
				query = "https://www.bing.com/search?q=bing"
			} else {
				query = "https://www.bing.com/search?q=" + remainingString
			}

			result := HandlerBing("GET", query)

			// processing

			response := processBingResponses(result)
			responseJSON := jsonResponseQuery {
				Status: true,
				Message: "here are the top search results",
				Result: response,
			}
			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech(responseJSON.Message, 0)

		} else if strings.ToLower(matchPars) == "youtube" || strings.ToLower(matchPars) == "videos" || strings.ToLower(matchPars) == "watch" {
			query := ""
			if len(remainingString) == 0 {
				query = "https://www.youtube.com/results?search_query=youtube"
			} else {
				query = "https://www.youtube.com/results?search_query=" + remainingString
			}

			result := HandlerYoutube("GET", query)

			// processing

			response := processYoutubeResponses(result)
			responseJSON := jsonResponseQuery {
				Status: true,
				Message: "here are the top search videos",
				Result: response,
			}
			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech(responseJSON.Message, 0)

		} else if strings.ToLower(matchPars) == "images" || strings.ToLower(matchPars) =="image"  {
			query := ""
			if len(remainingString) == 0 {
				query = "https://www.google.co.in/search?q="+"images"+"&source=lnms&tbm=isch"
			} else {
				query = "https://www.google.co.in/search?q="+remainingString+"&source=lnms&tbm=isch"
			}

			result := HandlerImage("GET", query)
			// processing

			response := processImageResponses(result)
			responseJSON := jsonResponseQuery {
				Status: true,
				Message: "here are the searched images",
				Result: response,
			}
			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech(responseJSON.Message, 0)

		} else if strings.ToLower(matchPars) == "weather" {

			if len(messageArr) == 1 || len(messageArr) < 3 {
				w.Write([]byte(`{"status": "success", "message": "ENTER: weather <city> <state>", "result": ""}`))
			} else {
				city := messageArr[len(messageArr)-2]
				state := messageArr[len(messageArr)-1]
				result := HandlerWeather(city, state)
				stringified, _ := json.Marshal(processWeather(result))
				response := jsonResponseWeather{
					Status: true,
					Message: "here are the current weather conditions",
					Result: string(stringified),
				}
				jData, _ := json.Marshal(response)
				w.Write(jData)
				TextToSpeech(response.Message + city + " " + state, 0)
			}
		} else if strings.ToLower(matchPars) == "meaning" {

			if len(messageArr) == 1 {
				w.Write([]byte(`{"status": "success", "message": "ENTER: meaning <word>", "result": ""}`))
			} else {

				wordStr := remainingString

				result := HandlerMeaning(wordStr)
				response := processMeaning(result)
				fmt.Println(len(response))

				if len(response) > 0 {

					responseJSON := jsonResponseMeaning{
					Status: true,
					Message: "here is the meaning of the searched word",
					Result: response,
					}

					jData, _ := json.Marshal(responseJSON)
					w.Write(jData)
					TextToSpeech(responseJSON.Message + " " + filterForSpeech(wordStr), 0)

				} else {
					query := "https://www.google.co.in/search?q=" + wordStr
					result := HandlerGoogle("GET", query)
					response := processGoogleResponses(result)

					responseJSON := jsonResponseQuery{
					Status: true,
					Message: "here are the top search results",
					Result: response,
					}

					jData, _ := json.Marshal(responseJSON)
					w.Write(jData)
					TextToSpeech(responseJSON.Message + " " + filterForSpeech(wordStr), 0)

				}
			}
		} else if strings.ToLower(matchPars) == "medicine" {

			if len(messageArr) <= 1 {
				w.Write([]byte(`{"status": "success", "message": "ENTER: medicine <generic / common name>", "result": ""}`))
			} else {
				med := messageArr[len(messageArr)-1]
				result := messages.HealthMedController(med, w)
				TextToSpeech(result, 0)
			}
		} else if strings.ToLower(matchPars) == "symptoms" {
			// add support for multiple symptoms at once and use ML to determine the best medicine suited

			if len(messageArr) < 2 {
				w.Write([]byte(`{"status": "success", "message": "ENTER: symptoms <symptom / condition>", "result": ""}`))
			} else {
				fmt.Println("inside")
				symp := strings.Join(messageArr[1:len(messageArr)], " ")
				result := messages.HealthSympController(symp, w)
				TextToSpeech(result, 0)
			}
		} else if strings.HasPrefix(strings.ToLower(message),"set reminder") {
			w.Write([]byte(`{"status": "success", "message": "Enter Reminder details : ", "result": ""}`))
		} else if strings.HasPrefix(strings.ToLower(message),"show reminder") {
			result := ShowReminder()
			fmt.Println(result)
			responseJSON := reminderResponse {
				Status: true,
				Message: "Here are your reminders : ",
				Result: result,
			}
			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech("Here are your reminders.", 0)
		} else if strings.HasPrefix(strings.ToLower(message),"deploy") {
			// support for deployment functionality
			fmt.Println("remaining string ", messageArr[len(messageArr) - 1])
			status := herokuhost.DeploymentFunction(messageArr[len(messageArr) - 1], w)
			TextToSpeech(filterForSpeech(status), 0)
		} else if strings.HasPrefix(strings.ToLower(message),"send mail") {
			w.Write([]byte(`{"status": "success", "message": "Enter Mail details : ", "result": ""}`))
		} else {
			// general conversation
			speech := messages.GeneralConvHandler(routeObject.message, routeObject.username, w)
			TextToSpeech(filterForSpeech(speech), 0)
		}
	} else {

		if strings.ToLower(matchPars) == "google" || strings.ToLower(matchPars) == "yahoo" || strings.ToLower(matchPars) == "bing" || strings.ToLower(matchPars) == "youtube" || 
			strings.ToLower(matchPars) == "image" || strings.ToLower(matchPars) == "weather" || strings.ToLower(matchPars) == "medicine" || strings.ToLower(matchPars) == "symptoms" ||
			strings.HasPrefix(strings.ToLower(message), "send mail"){

			w.Write([]byte(`{"status": "success", "message": "Services unavailable at the moment ! Check your Internet Connection and try again.", "result": ""}`))
			TextToSpeech("Services unavailable at the moment!", 0)

		} else if strings.HasPrefix(strings.ToLower(message),"set reminder") {
			w.Write([]byte(`{"status": "success", "message": "Enter Reminder details : ", "result": ""}`))
		} else if strings.HasPrefix(strings.ToLower(message),"show reminder") {
			result := ShowReminder()
			fmt.Println(result)
			responseJSON := reminderResponse {
				Status: true,
				Message: "Here are your reminders : ",
				Result: result,
			}
			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech("Here are your reminders.", 0)
		} else {
			// general conversation
			speech := messages.GeneralConvHandler(routeObject.message, routeObject.username, w)
			TextToSpeech(filterForSpeech(speech), 0)
		}
	}

}

// customSort() to sort an array according to the order defined by another array
func customSort(arr1 []string, arr2 []string, m,n int) []string{
	freq := make(map[string]int)

	for i := 0; i < m; i++ {
		freq[arr1[i]]++;
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

func filterForSpeech(s string) string {

	s = strings.Replace(s, "?", " ", -1)
	s = strings.Replace(s, "%", " ", -1)
	s = strings.Replace(s, "#", " ", -1)
	s = strings.Replace(s, "$", " ", -1)
	s = strings.Replace(s, "@", " ", -1)
	s = strings.Replace(s, "&", " ", -1)
	s = strings.Replace(s, "^", " ", -1)
	s = strings.Replace(s, "*", " ", -1)
	s = strings.Replace(s, "/", ", ", -1)
	return s

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

// processes google query result, scraps the required data and returns it
func processGoogleResponses(result string) []messageQueryBody {

	subsl := "<h3 class=\"LC20lb\">"
	lensubsl := len(subsl)
	subsl2 := "</h3>"
	lensubsl2 := len(subsl2)
	subsl3 := "<cite"
	lensubsl3 := len(subsl3)
	subsl4 := "</cite>"
	lensubsl4 := len(subsl4)
	subsl5 := "\"st\""
	lensubsl5 := len(subsl5)
	lenresult := len(result)

	var queryResult messageQueryBody
	var queryResultArray []messageQueryBody
	for i := 0; i < lenresult - lensubsl; i++ {
		mess := ""
		if result[i : i + lensubsl] == subsl {
			length := i + lensubsl
			var last int
			for j:=1; ; j++ {
				if result[length + j: length + j + lensubsl2] == subsl2 {
					mess = result[length+21: length + j-7]
					queryResult.Head = mess
					last = length + j + lensubsl2
					i = last
					break
				}
			}

			found := false
			for j:= 1; ; j++ {
				if result[last + j: last + j + lensubsl3] == subsl3 { // matched found for "<cite"
					for k:= 1; ; k++ {
						if result[last + j + lensubsl3 + k: last + j + lensubsl3 + k + 1] == ">" { // finding index for ">"
							mid := last + j + lensubsl3 + k + 1
							for l := 1; ; l++ {
								if result[mid + l: mid + l + lensubsl4] == subsl4 {
									last = mid + l + lensubsl4
									link := result[mid : mid +l]
									i = last 
									found = true
									if len(link) >= 7 {
										if link[0: 7] != "http://" && link[0: 8] != "https://" {
											link = "http://" + link
										}
									}
									queryResult.Link = link
									break
								}
							} 
							break
						}
					}

					for l := 1; i + l + lensubsl5 < len(result) ;l++ {
						if result[i + l: i + l + lensubsl5] == subsl5 {
							length = i + lensubsl5 + l + 1
							for m := 1; ; m++ {
								if result[length + m: length + m +7] == "</span>" {
									desc := result[length : length + m]
									queryResult.Desc = desc
									i = length + m + 6
									break;
								}
							}
							break;
						}
					}
				}
				if found {
					queryResultArray = append(queryResultArray, queryResult)
					break
				}
			}
		}
	}
	return queryResultArray

}

func processWeather(response string) weatherStr  {

	fmt.Println("this is the response")
	fmt.Println(response)
	subl := "in json format"
	sublLen := len(subl)
	found := false
	var weather []byte
	var weatherInJSON weatherStr
	for i:=0; i< len(response) - sublLen; i++ {
		if response[i: i + sublLen] == subl {
			for j:=1; ; j++ {
				if response[i+sublLen+j: i+sublLen+j + 1] == "}" {
					weather = []byte(response[i+sublLen+1: i+sublLen+j+1])
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	if !found {
		fmt.Println("corrupted logging!")
	}
	fmt.Println(string(weather))
	err := json.Unmarshal(weather, &weatherInJSON)
	if err != nil {
		panic(err)
	}
	fmt.Println(weatherInJSON)
	return weatherInJSON

}

// processes yahoo query result, scraps the required data and returns it
func processYahooResponses(result string) []messageQueryBody {

	subsl := "<a class=\" ac-algo fz-l ac-21th lh-24\"";
	lensubsl := len(subsl)
	subsl2 := "</a>"
	lensubsl2 := len(subsl2)
	subsl3 := "<span class=\" fz-ms fw-m fc-12th wr-bw lh-17\">"
	lensubsl3 := len(subsl3)
	subsl4 := "</span>"
	lensubsl4 := len(subsl4)
	subsl5 := "<p class=\"lh-16\""
	lensubsl5 := len(subsl5)

	var queryResult messageQueryBody
	var queryResultArray []messageQueryBody
	for i := 0; i < len(result) - lensubsl; i++ {
		mess := ""
		if result[i : i + lensubsl] == subsl {
			length := i + lensubsl
			var last int
			var start int

			for k := 1; ; k++ {
				if result[length + k: length+k+1 ] == ">" {
					start =  length + k + 1;
					break;
				}
			}

			for j:=1; ; j++ {
				if result[start + j: start + j + lensubsl2] == subsl2 {
					mess = result[start: start + j]
					queryResult.Head = mess
					last = start + j + lensubsl2
					i = last
					break
				}
			}

			found := false
			for j:= 1; ; j++ {
				if result[last + j: last + j + lensubsl3] == subsl3 { // matched found for "<span class=\" fz-ms fw-m fc-12th wr-bw lh-17\">"
					for k:= 1; ; k++ {
						if result[last + j + lensubsl3 + k: last + j + lensubsl3 + k + lensubsl4] == subsl4 { // finding index for "</span>"
							link := result[last + j + lensubsl3 : last + j + lensubsl3 + k]
							i = last + j + lensubsl3 + k + lensubsl4
							found = true
							link = strings.Replace(link, "<b>", "", -1)
							link = strings.Replace(link, "</b>", "", -1)
							if len(link) >= 7 {
								if link[0: 7] != "http://" && link[0: 8] != "https://" {
									link = "http://" + link
								}
							}
							queryResult.Link = link
							break
						}
					}
					for k := 1; ; k++ {
						if result[i + k : i + k + lensubsl5] == subsl5 {
							length = i + k + lensubsl5 + 1;
							for l := 1; ; l++ {
								if result[length + l: length + l + 4] == "</p>" {
									desc := result[length: length + l]
									queryResult.Desc = desc;
									i = length + l +4;
									break;
								}
							}
							break;
						}
					}
				}

				if found {
					queryResultArray = append(queryResultArray, queryResult)
					break
				}
			}
		}
	}
	return queryResultArray

}

// processes bing query result, scraps the required data and returns it
func processBingResponses(result string) []messageQueryBody {

	subsl := "<li class=\"b_algo\""
	subsl2 := "<a"
	subsl3 := "<cite"
	lensubsl3 := len(subsl3)
	subsl4 := "</cite>"
	lensubsl4 := len(subsl4)
	subsl5 := "<p>"
	lensubsl5 := len(subsl5)

	var queryResult messageQueryBody
	var queryResultArray []messageQueryBody

	for i := 0; i < len(result) - len(subsl); i++ {
		mess := ""
		if result[i : i + len(subsl)] == subsl {
			length := i + len(subsl)
			var last int
			var aStart int
			var start int

			for k := 1; ; k++ {
				if result[length + k: length + k + 2 ] == subsl2 {
					aStart = length + k
					for l := 1; ; l++ {
						if result[aStart + l: aStart + l + 1 ] == ">" {
							start = aStart + l + 1;
							break;
						}
					}
					break;
				}
			}

			for j:=1; ; j++ {
				if result[start + j: start + j + 4] == "</a>" {
					mess = result[start: start + j]
					queryResult.Head = mess
					last = start + j + 4
					i = last
					break
				}
			}

			found := false
			for j:= 1; ; j++ {
				if result[last + j: last + j + lensubsl3] == subsl3 { // matched found for "<cite"
					for k:= 1; ; k++ {
						if result[last + j + lensubsl3 + k: last + j + lensubsl3 + k + lensubsl4] == subsl4 { // finding index for "</cite>"
							link := result[last + j + lensubsl3 + 1 : last + j + lensubsl3 + k]

							i = last + j + lensubsl3 + k + lensubsl4
							found = true
							link = strings.Replace(link, "<strong>", "", -1)
							link = strings.Replace(link, "</strong>", "", -1)
							if len(link) >= 7 {
								if link[0: 7] != "http://" && link[0: 8] != "https://" {
									link = "http://" + link
								}
							}
							
							queryResult.Link = link
							break
						}
					}
					for k := 1; ; k++ {
						if result[i + k : i + k + lensubsl5] == subsl5 {
							length = i + k + lensubsl5;
							for l := 1; ; l++ {
								if result[length + l: length + l + 4] == "</p>" {
									desc := result[length: length + l]
									queryResult.Desc = desc;
									i = length + l +4;
									break;
								}
							}
							break;
						}
					}
				}

				if found {
					queryResultArray = append(queryResultArray, queryResult)
					break
				}
			}
		}
	}
	return queryResultArray
}

// processes youtube query result, scraps the required data and returns it
func processYoutubeResponses(result string) []messageQueryBody {

	subsl := "<a id=\"video-title\""
	subsl2 := "href=\""
	subsl3 := "</a>"
	lensubsl3 := len(subsl3)
	subsl4 := "<yt-formatted-string id=\"description-text\" class=\"style-scope ytd-video-renderer\">"
	lensubsl4 := len(subsl4)
	subsl5 := "</yt-formatted-string>"
	lensubsl5 := len(subsl5)

	var queryResult messageQueryBody
	var queryResultArray []messageQueryBody
	var mid int

	for i := 0; i < len(result) - len(subsl); i++ {
		mess := ""
		if result[i : i + len(subsl)] == subsl {
			length := i + len(subsl)
			var last int
			for j:=1; ; j++ {
				if result[length + j: length + j + len(subsl2)] == subsl2 {
					mid = length + j + len(subsl2)
					for k := 1; ; k++ {
						if result[mid + k: mid + k + 1] == "\"" {
							link := result[mid: mid + k]
							flink := "https://www.youtube.com" + link
							queryResult.Link = flink
							last = mid + k + 1
							for l := 1; ; l++ {
								if result[last + l: last+ l +2] == "\">" {
									last = last + l +2
									i = last + l + 2
									break;
								}
							}
							break
						}
					}
					break
				}
			}

			found := false
			for j:= 1; ; j++ {
				if result[last + j: last + j + lensubsl3] == subsl3 { // matched found for "</a>"
					mess = result[last: last + j]
					i = last + j + lensubsl3
					found = true
					queryResult.Head = mess
					for k := 1; ; k++ {
						if result[i + k : i + k + lensubsl4] == subsl4 {
							length = i + k + lensubsl4;
							for l := 1; ; l++ {
								if result[length + l: length + l + lensubsl5] == subsl5 {
									desc := result[length: length + l]
									queryResult.Desc = desc;
									i = length + l +4;
									break;
								}
							}
							break;
						}
					}
				}
				if found {
					queryResultArray = append(queryResultArray, queryResult)
					break
				}
			}
		}
	}
	return queryResultArray

}

// processes image query result, scraps the required data and returns it
func processImageResponses(result string) []messageQueryBody {

	subsl := "<div class=\"rg_meta notranslate\">"
	lensubsl := len(subsl)
	subsl2 := "\"ou\":\""
	lensubsl2 := len(subsl2)
	subsl3 := "\"pt\":"
	lensubsl3 := len(subsl3)
	subsl4 := "\"rh\":"
	lensubsl4 := len(subsl4)
	count := 0

	var queryResult messageQueryBody
	var queryResultArray []messageQueryBody

	for i := 0; i < len(result) - len(subsl); i++ {
		link := ""
		if result[i : i + lensubsl] == subsl {
			length := i + lensubsl
			var mid int
			for j := 1; ; j++ {
				found := false
				if result[length + j: length + j + lensubsl2] == subsl2 {
					mid = length + j + lensubsl2
					for k := 1; ; k++ {
						if result[mid + k: mid + k + 1] == "\"" {
							link = result[mid: mid + k]
							queryResult.Link = link
							found = true
							i = mid + k + 1;
							break;
						}
					}

					for a := 1; ; a++ {
						if result[i + a: i + a + lensubsl3] == subsl3 {
							mid = i + a + lensubsl3 + 1
							for k := 1; ; k++ {
								if result[mid + k: mid + k + 1] == "\"" {
									desc := result[mid: mid + k]
									queryResult.Desc = desc
									found = true
									i = mid + k + 1;
									break;
								}
							}
							break;
						}
					}

					for a := 1; ; a++ {
						if result[i + a: i + a + lensubsl4] == subsl4 {
							mid = i + a + lensubsl4 + 1
							for k := 1; ; k++ {
								if result[mid + k: mid + k + 1] == "\"" {
									dlink := result[mid: mid + k]
									queryResult.DescLink = dlink
									found = true
									i = mid + k + 1;
									break;
								}
							}
							break;
						}
					}
				}
				if found {
					queryResultArray = append(queryResultArray, queryResult)
					count ++
					break;
				}
			}
		}
		if count == 10 {
			break;
		}
	}
	return queryResultArray

}

func processMeaning(response string) []meaningStr {

	fmt.Println("this is the response")

	found := false
	
	subs1, subs2, subs3, subs4  := "meaning", "example", "subMeaning", "subExample"
	subsLen1, subsLen2, subsLen3, subsLen4 := len(subs1), len(subs2), len(subs3), len(subs4)
	var mid, last = 0, 0

	var meaningBody meaningStr
	var meaningBodyArray []meaningStr
	var subMeaningBody submeanStr
	var subMeaningBodyArray []submeanStr

	for i:=0; i< len(response) - subsLen1; i++ {
		found = false
		if response[i: i + subsLen1] == subs1 {
			subMeaningBodyArray = nil
			mid = i + subsLen1 + 4
			last = mid
			for j:=1; j< len(response) - last - 1 ; j++ {
				if response[last+j: last + j + 1] == "*" {
					meaningBody.Meaning = response[mid: mid + j]
					found = true
					last = mid + j +1

					for k:=1; k < len(response) - last - subsLen2 ; k++ {
						if response[last + k: last + k + subsLen2] == subs2 {
							v := mid + j + k + subsLen2 + 4
							last = v
							for l:= 1; l < len(response) - last - 1; l++ {
								if response[v + l: v + l + 1] == "*" {
									meaningBody.Example = response[v + 1 : v + l]
									last = v + l +1
									break;
								}
							}
							break;
						}
					}
					
					for k := 1;k < len(response) - last -subsLen3 ; k++ {
						check := false
						if response[last + k: last + k + subsLen3] == subs3 {
							v := last + k + subsLen3 + 4
							last := v
							for l:= 1; l < len(response) - last - 1; l++ {
								if response[v +l: v + l + 1] == "*" {
									subMeaningBody.Smean = response[v : v + l]
									last = v + l + 1
									check = true
									break;
								}
							}
							if check {
								for m:= 1; m < len(response) - last - subsLen4 ; m++ {
									if response[last + m: last + m + subsLen4] == subs4 {
										v := last + m + subsLen4 + 4
										last = v
										for l:= 1; l < len(response) - last - 1 ; l++ {
											if response[v +l: v + l + 1] == "*" {
												subMeaningBody.Subexample = response[v : v + l]
												last = v + l + 1
												break;
											}
										}
										subMeaningBodyArray = append(subMeaningBodyArray, subMeaningBody)
										meaningBody.Submeaning = subMeaningBodyArray
										break;
									}
								}
							}
							break;
						}
					}
					break;	
				}
			}
			if found {
				meaningBodyArray = append(meaningBodyArray, meaningBody)
				i = last
			}
		}
	}
	fmt.Println("meaningBodyArray : ", meaningBodyArray)
	
	return meaningBodyArray

}