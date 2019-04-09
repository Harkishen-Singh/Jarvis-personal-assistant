package controllers

import (
	"net/http"
	"strings"
	"encoding/json"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/messages"
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
}

// type imageQueryBody struct{
// 	Link string `json:"link"`
// }

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

	/*message := routeObject.message
	messageArr := strings.Split(message, " ")
	// messageTemp := message
	var firstPars string
	if strings.Contains(message, " ") {
		firstPars = message[:strings.Index(message, " ")]
	} else {
		firstPars = message
	}

	strArr := strings.Split(firstPars, " ")
	strArrDiff := strings.Split(message, " ")

	remainingString := strings.Join(stringDifference(strArr, strArrDiff), " ")*/
	// lastParsArr := strings.Split(messageTemp, " ")
	// lastPars := lastParsArr[len(lastParsArr) - 1]
	message := routeObject.message
	messageArr := strings.Split(message, " ")
	var firstPars string
	var secondPars string
	var thirdPars string
	if len(messageArr)>=3 {
		firstPars=messageArr[0]
		secondPars = messageArr[1]
		thirdPars = messageArr[2]
	} else if len(messageArr)==2 {
		firstPars=messageArr[0]
		secondPars = messageArr[1]
	} else if(len(messageArr)==1) {
			firstPars=messageArr[0]
	}

	var remainingString string
	var strArr []string
	var strArr0 string
	var strArrDiff []string
	strArrDiff = strings.Split(message, " ")
	if strings.ToLower(firstPars) == "google" || strings.ToLower(firstPars) =="yahoo" || strings.ToLower(firstPars) =="bing" || strings.ToLower(firstPars) =="search" || strings.ToLower(firstPars) =="youtube"|| strings.ToLower(firstPars) == "watch"||strings.ToLower(firstPars) =="videos" || strings.ToLower(firstPars) =="images" || strings.ToLower(firstPars) =="image" {
		if strings.ToLower(secondPars) == "search" || strings.ToLower(secondPars) =="google" ||strings.ToLower(secondPars) =="yahoo" ||strings.ToLower(secondPars) =="bing" ||strings.ToLower(secondPars) =="videos" ||strings.ToLower(secondPars) =="youtube"||strings.ToLower(secondPars) =="for" || strings.ToLower(secondPars) =="of"{
			if strings.ToLower(thirdPars) == "for" || strings.ToLower(thirdPars) =="videos" {
				strArr0=firstPars+" "+secondPars+" "+thirdPars
				strArr=strings.Split(strArr0," ")
				remainingString=strings.Join(stringDifference(strArr,strArrDiff)," ")
				fmt.Println(remainingString)
			} else {
				strArr0=firstPars+" "+secondPars
				strArr=strings.Split(strArr0," ")
				remainingString=strings.Join(stringDifference(strArr,strArrDiff)," ")
				fmt.Println(remainingString)
			}
		} else {
				strArr = strings.Split(firstPars, " ")
				strArrDiff = strings.Split(message, " ")
				remainingString=strings.Join(stringDifference(strArr,strArrDiff)," ")
				fmt.Println(remainingString)
			}
	}
	// single word operations


	if Connected() {

		if strings.ToLower(firstPars) == "google" || strings.ToLower(secondPars) == "google" { // for google search
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

		} else if strings.ToLower(firstPars) == "yahoo" || strings.ToLower(secondPars) == "yahoo"{
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

		} else if strings.ToLower(firstPars) == "bing" || strings.ToLower(secondPars)=="bing"{
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

		} else if strings.ToLower(firstPars) == "youtube" || strings.ToLower(firstPars) == "videos" || strings.ToLower(firstPars) == "watch" || strings.ToLower(secondPars) == "youtube"{
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
				Message: "here are the top search results",
				Result: response,
			}
			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech(responseJSON.Message, 0)

		} else if strings.ToLower(firstPars) == "images" || strings.ToLower(firstPars) =="image"  {
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

		} else if strings.ToLower(firstPars) == "weather" {

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
		} else {
			// general conversation
			speech := messages.GeneralConvHandler(routeObject.message, routeObject.username, w)
			TextToSpeech(filterForSpeech(speech), 0)
		}
	} else {

		if strings.ToLower(firstPars) == "google" || strings.ToLower(firstPars) == "yahoo" || strings.ToLower(firstPars) == "bing" || strings.ToLower(firstPars) == "youtube" || strings.ToLower(firstPars) == "image" || strings.ToLower(firstPars) == "weather" {
			w.Write([]byte(`{"status": "success", "message": "Services unavailable at the moment ! Check your Internet Connection and try again.", "result": ""}`))
			TextToSpeech("Services unavailable at the moment!", 0)
		} else {
			// general conversation
			speech := messages.GeneralConvHandler(routeObject.message, routeObject.username, w)
			TextToSpeech(filterForSpeech(speech), 0)
		}
	}

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
	var queryResult messageQueryBody
	var queryResultArray []messageQueryBody
	for i := 0; i < len(result) - lensubsl; i++ {
		mess := ""
		if result[i : i + lensubsl] == subsl {
			length := i + lensubsl
			var last int
			for j:=1; ; j++ {
				if result[length + j: length + j + lensubsl2] == subsl2 {
					mess = result[length: length + j]
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
						if result[last + j + lensubsl3 + k: last + j + lensubsl3 + k + lensubsl4] == subsl4 { // finding index for "</cite>"
							link := result[last + j + lensubsl3 + 15 : last + j + lensubsl3 + k]
							i = last + j + lensubsl3 + k + lensubsl4
							found = true
							if link[0: 7] != "http://" &&  link[0: 4] != "www." && link[0: 8] != "https://" {
								link = "http://" + link
							}
							queryResult.Link = link
							break
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
							if link[0: 7] != "http://" &&  link[0: 4] != "www." {
								link = "http://" + link
							}
							queryResult.Link = link
							break
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
					fMess := strings.Replace(mess, "<strong>", "", -1)
					finalMess := strings.Replace(fMess, "</strong>", "", -1)
					queryResult.Head = finalMess
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
							if link[0: 7] != "http://" &&  link[0: 4] != "www." {
								link = "http://" + link
							}
							queryResult.Link = link
							break
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
						if result[mid + k: mid + k + 2] == "\">" {
							link := result[mid: mid + k]
							flink := "https://www.youtube.com" + link
							queryResult.Link = flink
							last = mid + k + 2
							i = last
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
