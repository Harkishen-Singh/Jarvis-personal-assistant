package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/config"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/logger"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/messages"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/scrapper"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/services/herokuhost"
	"github.com/PuerkitoBio/goquery"
)

type response struct {
	username string
	message  string
}

type statusCode struct {
	status string
}

type messageQueryBody struct {
	Head     string `json:"head"`
	Link     string `json:"link"`
	Desc     string `json:"desc"`
	DescLink string `json:"dlink"`
}

type reminderResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Result  []reminder `json:"result"`
}

type jsonResponseQuery struct {
	Status  bool               `json:"status"`
	Message string             `json:"message"`
	Result  []messageQueryBody `json:"result"`
}

type jsonResponseWeather struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Result  weatherStr `json:"result"`
}

type weatherStr struct {
	Time        string `json:"time"`
	City        string `json:"city"`
	Temperature string `json:"temperature"`
	DewPoint    string `json:"dew_point"`
	Humidity    string `json:"humidity"`
	Visibility  string `json:"visibility"`
	FeelsLike   string `json:"feels_like"`
}

type meaningStr struct {
	Meaning    string       `json:"meaning"`
	Example    string       `json:"example"`
	Submeaning []submeanStr `json:"submeaning"`
}

type submeanStr struct {
	Smean      string
	Subexample string
}

type jsonResponseMeaning struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Result  []meaningStr `json:"result"`
}

// MessagesController controls messages handling
func MessagesController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()

	request := response{
		username: r.FormValue("username"),
		message:  r.FormValue("message"),
	}
	fmt.Println(request)

	routes(request, w)

}

func routes(routeObject response, w http.ResponseWriter) {

	message := routeObject.message
	message = strings.ToLower(message)
	messageArr := strings.Fields(message)
	var a []string
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
	matchPars := ""
	remainingString := ""
	if len(a) > 0 {
		sort := customSort(a, priority, len(a), len(priority))
		matchPars = sort[0]
		remainingString = strings.Join(messageArr[:], " ")
		messageArr = append([]string{matchPars}, messageArr...)
	} else {
		remainingString = strings.Join(messageArr[:], " ")
	}

	if Connected() {

		if strings.ToLower(matchPars) == "google" || strings.ToLower(matchPars) == "search" { // for google search
			if len(remainingString) == 0 {
				remainingString = "google"
			}

			// processing

			response, _ := processGoogleResponses(remainingString, "com", "en", nil, 1, 10, 5)
			responseJSON := jsonResponseQuery{
				Status:  true,
				Message: "here are the top search results",
				Result:  response,
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
			responseJSON := jsonResponseQuery{
				Status:  true,
				Message: "here are the top search results",
				Result:  response,
			}
			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech(responseJSON.Message, 0)

		} else if strings.ToLower(matchPars) == "bing" {
			if len(remainingString) == 0 {
				remainingString = "bing"
			}

			// processing

			response, _ := processBingResponses(remainingString, "com", nil, 1, 10, 5)
			responseJSON := jsonResponseQuery{
				Status:  true,
				Message: "here are the top search results",
				Result:  response,
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
			responseJSON := jsonResponseQuery{
				Status:  true,
				Message: "here are the top search videos",
				Result:  response,
			}
			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech(responseJSON.Message, 0)

		} else if strings.ToLower(matchPars) == "images" || strings.ToLower(matchPars) == "image" {
			query := ""
			if len(remainingString) == 0 {
				query = "images"
			} else {
				query = remainingString
			}

			result := scrapeImage(query)

			responseJSON := jsonResponseQuery{
				Status:  true,
				Message: "here are the searched images",
				Result:  result,
			}
			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech(responseJSON.Message, 0)

		} else if strings.ToLower(matchPars) == "weather" {

			if len(messageArr) < 3 {
				w.Write([]byte(`{"status": "success", "message": "ENTER: weather <city> <state>", "result": ""}`))
			} else {
				city := messageArr[len(messageArr)-2]
				state := messageArr[len(messageArr)-1]
				weatherData := getWeather(city, state)
				response := jsonResponseWeather{
					Status:  true,
					Message: "here are the current weather conditions",
					Result:  weatherData,
				}
				jData, _ := json.Marshal(response)
				w.Write(jData)
				TextToSpeech(response.Message+city+" "+state, 0)
			}
		} else if strings.ToLower(matchPars) == "meaning" {

			if len(messageArr) == 1 {
				w.Write([]byte(`{"status": "success", "message": "ENTER: meaning <word>", "result": ""}`))
			} else {

				wordStr := remainingString

				response := getMeaning(wordStr)
				fmt.Println(len(response))

				if len(response) > 0 {

					responseJSON := jsonResponseMeaning{
						Status:  true,
						Message: "here is the meaning of the searched word",
						Result:  response,
					}

					jData, _ := json.Marshal(responseJSON)
					w.Write(jData)
					TextToSpeech(responseJSON.Message+" "+filterForSpeech(wordStr), 0)

				} else {
					response, _ := processGoogleResponses(wordStr, "com", "en", nil, 1, 10, 5)

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
		} else if strings.HasPrefix(strings.ToLower(message), "set reminder") {
			w.Write([]byte(`{"status": "success", "message": "Enter Reminder details : ", "result": ""}`))
		} else if strings.HasPrefix(strings.ToLower(message), "show reminder") {
			result := ShowReminder()
			fmt.Println(result)
			responseJSON := reminderResponse{
				Status:  true,
				Message: "Here are your reminders : ",
				Result:  result,
			}
			jData, _ := json.Marshal(responseJSON)
			w.Write(jData)
			TextToSpeech("Here are your reminders.", 0)
		} else if strings.HasPrefix(strings.ToLower(message), "deploy") {
			// support for deployment functionality
			fmt.Println("remaining string ", messageArr[len(messageArr)-1])
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

		if strings.ToLower(matchPars) == "google" || strings.ToLower(matchPars) == "yahoo" || strings.ToLower(matchPars) == "bing" || strings.ToLower(matchPars) == "youtube" ||
			strings.ToLower(matchPars) == "image" || strings.ToLower(matchPars) == "weather" || strings.ToLower(matchPars) == "medicine" || strings.ToLower(matchPars) == "symptoms" ||
			strings.HasPrefix(strings.ToLower(message), "send mail") {

			w.Write([]byte(`{"status": "success", "message": "Services unavailable at the moment ! Check your Internet Connection and try again.", "result": ""}`))
			TextToSpeech("Services unavailable at the moment!", 0)

		} else if strings.HasPrefix(strings.ToLower(message), "set reminder") {
			w.Write([]byte(`{"status": "success", "message": "Enter Reminder details : ", "result": ""}`))
		} else if strings.HasPrefix(strings.ToLower(message), "show reminder") {
			result := ShowReminder()
			fmt.Println(result)
			responseJSON := reminderResponse{
				Status:  true,
				Message: "Here are your reminders : ",
				Result:  result,
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

func buildGoogleUrls(searchTerm, countryCode, languageCode string, pages, count int) ([]string, error) {
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := googleDomains[countryCode]; found {
		for i := 0; i < pages; i++ {
			start := i * count
			scrapeURL := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0", googleBase, searchTerm, count, languageCode, start)
			toScrape = append(toScrape, scrapeURL)
		}
	} else {
		err := fmt.Errorf("country (%s) is currently not supported", countryCode)
		return nil, err
	}
	return toScrape, nil
}

func googleResultParsing(response *http.Response, rank int) ([]messageQueryBody, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}
	results := []messageQueryBody{}
	sel := doc.Find("div.g")
	rank++
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		descTag := item.Find("span.st")
		desc := descTag.Text()
		link = strings.Trim(link, " ")
		if desc == "" {
			desc = link
		}
		if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
			result := messageQueryBody{
				Head: desc,
				Link: link,
			}
			results = append(results, result)
			rank++
		}
	}
	return results, err
}

// processGoogleResponses scrapes the relevant Google search engine for SearchResults
func processGoogleResponses(searchTerm, countryCode, languageCode string, proxyString interface{}, pages, count, backoff int) ([]messageQueryBody, error) {
	results := []messageQueryBody{}
	resultCounter := 0
	googlePages, err := buildGoogleUrls(searchTerm, countryCode, languageCode, pages, count)
	if err != nil {
		return nil, err
	}
	for _, page := range googlePages {
		res, err := scrapper.ScrapeClientRequest(page, proxyString)
		if err != nil {
			return nil, err
		}
		data, err := googleResultParsing(res, resultCounter)
		if err != nil {
			return nil, err
		}
		resultCounter += len(data)
		for _, result := range data {
			results = append(results, result)
		}
		time.Sleep(time.Duration(backoff) * time.Second)
	}
	return results, nil
}

// processes yahoo query result, scraps the required data and returns it
func processYahooResponses(result string) []messageQueryBody {

	subsl := "<a class=\" ac-algo fz-l ac-21th lh-24\""
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
	for i := 0; i < len(result)-lensubsl; i++ {
		mess := ""
		if result[i:i+lensubsl] == subsl {
			length := i + lensubsl
			var last int
			var start int

			for k := 1; ; k++ {
				if result[length+k:length+k+1] == ">" {
					start = length + k + 1
					break
				}
			}

			for j := 1; ; j++ {
				if result[start+j:start+j+lensubsl2] == subsl2 {
					mess = result[start : start+j]
					queryResult.Head = mess
					last = start + j + lensubsl2
					i = last
					break
				}
			}

			found := false
			for j := 1; ; j++ {
				if result[last+j:last+j+lensubsl3] == subsl3 { // matched found for "<span class=\" fz-ms fw-m fc-12th wr-bw lh-17\">"
					for k := 1; ; k++ {
						if result[last+j+lensubsl3+k:last+j+lensubsl3+k+lensubsl4] == subsl4 { // finding index for "</span>"
							link := result[last+j+lensubsl3 : last+j+lensubsl3+k]
							i = last + j + lensubsl3 + k + lensubsl4
							found = true
							link = strings.Replace(link, "<b>", "", -1)
							link = strings.Replace(link, "</b>", "", -1)
							if len(link) >= 7 {
								if link[0:7] != "http://" && link[0:8] != "https://" {
									link = "http://" + link
								}
							}
							queryResult.Link = link
							break
						}
					}
					for k := 1; ; k++ {
						if result[i+k:i+k+lensubsl5] == subsl5 {
							length = i + k + lensubsl5 + 1
							for l := 1; ; l++ {
								if result[length+l:length+l+4] == "</p>" {
									desc := result[length : length+l]
									queryResult.Desc = desc
									i = length + l + 4
									break
								}
							}
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

func buildBingUrls(searchTerm, country string, pages, count int) ([]string, error) {
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if countryCode, found := bingDomains[country]; found {
		for i := 0; i < pages; i++ {
			first := firstParameter(i, count)
			scrapeURL := fmt.Sprintf("https://bing.com/search?q=%s&first=%d&count=%d%s", searchTerm, first, count, countryCode)
			toScrape = append(toScrape, scrapeURL)
		}
	} else {
		err := fmt.Errorf("country (%s) is currently not supported", country)
		return nil, err
	}
	return toScrape, nil
}

func firstParameter(number, count int) int {
	if number == 0 {
		return number + 1
	}
	return number*count + 1
}

func bingResultParser(response *http.Response, rank int) ([]messageQueryBody, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}
	results := []messageQueryBody{}
	sel := doc.Find("li.b_algo")
	rank++
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		descTag := item.Find("div.b_caption p")
		desc := descTag.Text()
		link = strings.Trim(link, " ")
		if desc == "" {
			desc = link
		}
		if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
			result := messageQueryBody{
				Head: desc,
				Link: link,
			}
			results = append(results, result)
			rank++
		}
	}
	return results, err
}

// processBingResponses scrapes bing.com with desired parameters
func processBingResponses(searchTerm, country string, proxyString interface{}, pages, count, backoff int) ([]messageQueryBody, error) {
	results := []messageQueryBody{}
	bingPages, err := buildBingUrls(searchTerm, country, pages, count)
	if err != nil {
		return nil, err
	}
	for _, page := range bingPages {
		rank := len(results)
		fmt.Println("page:: ", page)
		res, err := scrapper.ScrapeClientRequest(page, proxyString)
		if err != nil {
			return nil, err
		}
		data, err := bingResultParser(res, rank)
		if err != nil {
			return nil, err
		}
		for _, result := range data {
			results = append(results, result)
		}
		time.Sleep(time.Duration(backoff) * time.Second)
	}
	return results, nil
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

	for i := 0; i < len(result)-len(subsl); i++ {
		mess := ""
		if result[i:i+len(subsl)] == subsl {
			length := i + len(subsl)
			var last int
			for j := 1; ; j++ {
				if result[length+j:length+j+len(subsl2)] == subsl2 {
					mid = length + j + len(subsl2)
					for k := 1; ; k++ {
						if result[mid+k:mid+k+1] == "\"" {
							link := result[mid : mid+k]
							flink := "https://www.youtube.com" + link
							queryResult.Link = flink
							last = mid + k + 1
							for l := 1; ; l++ {
								if result[last+l:last+l+2] == "\">" {
									last = last + l + 2
									i = last + l + 2
									break
								}
							}
							break
						}
					}
					break
				}
			}

			found := false
			for j := 1; ; j++ {
				if result[last+j:last+j+lensubsl3] == subsl3 { // matched found for "</a>"
					mess = result[last : last+j]
					i = last + j + lensubsl3
					found = true
					queryResult.Head = mess
					for k := 1; ; k++ {
						if result[i+k:i+k+lensubsl4] == subsl4 {
							length = i + k + lensubsl4
							for l := 1; ; l++ {
								if result[length+l:length+l+lensubsl5] == subsl5 {
									desc := result[length : length+l]
									queryResult.Desc = desc
									i = length + l + 4
									break
								}
							}
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

func getMeaning(word string) []meaningStr {
	url := "https://en.oxforddictionaries.com/definition/" + word
	res, err := scrapper.ScrapeClientRequest(url, nil)
	if err != nil {
		logger.Error(err)
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		logger.Error(err)
	}
	var resultObj []meaningStr
	sel := doc.Find("section.gramb")
	for sectionIndex := range sel.Nodes {
		fmt.Println(sectionIndex)
		section := sel.Eq(sectionIndex)
		typArray := section.Find("span.pos")
		for typIndex := range typArray.Nodes {
			fmt.Println(typArray.Eq(typIndex).Text())
		}
		trgArray := section.Find("div.trg")
		for trgIndex := range trgArray.Nodes {
			var meaning meaningStr
			trg := trgArray.Eq(trgIndex)
			paraArray := trg.Find("p")
			for paraIndex := range paraArray.Nodes {
				para := paraArray.Eq(paraIndex)
				meaningArray := para.Find("span.ind")
				for meaningIndex := range meaningArray.Nodes {
					meaning.Meaning = (meaningArray.Eq(meaningIndex).Text())
				}
			}

			exampleArray := trg.Find("li.ex")
			if len(exampleArray.Nodes) > 0 {
				meaning.Example = exampleArray.Eq(0).Text()
			} else {
				exampleArray := trg.Find("div.ex")
				if len(exampleArray.Nodes) > 0 {
					meaning.Example = exampleArray.Eq(0).Text()
				}
			}

			subMeaningArray := trg.Find("li.subSense")
			var extractedSubMean []submeanStr
			for subMeaningIndex := range subMeaningArray.Nodes {
				var subMeaningObj submeanStr
				subMeaningBlock := subMeaningArray.Eq(subMeaningIndex)
				meaningArray := subMeaningBlock.Find("span.ind")
				if len(meaningArray.Nodes) > 0 {
					subMeaningObj.Smean = meaningArray.Eq(0).Text()
				}

				exampleArray := subMeaningBlock.Find("li.ex")
				if len(exampleArray.Nodes) > 0 {
					subMeaningObj.Subexample = exampleArray.Eq(0).Text()
				} else {
					exampleArray := subMeaningBlock.Find("div.ex")
					if len(exampleArray.Nodes) > 0 {
						subMeaningObj.Subexample = exampleArray.Eq(0).Text()
					}
				}

				extractedSubMean = append(extractedSubMean, subMeaningObj)
			}
			meaning.Submeaning = extractedSubMean
			if meaning.Meaning != "" {
				resultObj = append(resultObj, meaning)
			}
		}
	}
	return resultObj
}

func getWeather(city string, state string) weatherStr {
	var resultObj weatherStr
	resultObj.City = city
	resultObj.Time = time.Now().Format("02-01-2006 3:4 PM")

	country := "india"
	url := "https://www.msn.com/en-in/weather/today/" + city + "," + state + "," + country + "/we-city?weadegreetype=C"
	res, err := scrapper.ScrapeClientRequest(url, nil)
	if err != nil {
		logger.Error(err)
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		logger.Error(err)
	}
	currentSectionArray := doc.Find("section.curcond")
	if len(currentSectionArray.Nodes) <= 0 {
		logger.Error(err)
	}
	currentSection := currentSectionArray.Eq(0)
	tempSpanArray := currentSection.Find("span.current")
	if len(tempSpanArray.Nodes) <= 0 {
		logger.Error(err)
	}
	tempSpan := tempSpanArray.Eq(0)
	resultObj.Temperature = tempSpan.Text() + "°C"

	locationArray := doc.Find("div.mylocations > div.header > span")
	if len(locationArray.Nodes) > 0 {
		resultObj.City = locationArray.Eq(0).Text()
	}

	conditionsArray := currentSection.Find("div.weather-info>ul>li")
	for conditionIndex := range conditionsArray.Nodes {
		conditionSection := conditionsArray.Eq(conditionIndex)
		conditionChildrens := conditionSection.Contents()
		if len(conditionChildrens.Nodes) <= 0 {
			continue
		}
		conditionName := strings.TrimSpace(conditionChildrens.First().Text())
		conditionVal := strings.TrimSpace(conditionChildrens.Last().Text())
		switch strings.ToLower(conditionName) {
		case "feels like":
			resultObj.FeelsLike = conditionVal + "C"
		case "humidity":
			resultObj.Humidity = conditionVal
		case "dew point":
			resultObj.DewPoint = conditionVal
		case "visibility":
			resultObj.Visibility = conditionVal
		}
	}
	return resultObj
}

func scrapeImage(query string) []messageQueryBody {
	url := "https://www.google.co.in/search?q=" + query + "&source=lnms&tbm=isch&gbv=1"
	res, err := scrapper.ScrapeClientRequest(url, nil)
	if err != nil {
		logger.Error(err)
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		logger.Error(err)
	}
	var resultObj []messageQueryBody
	tableArray := doc.Find("table")
	if len(tableArray.Nodes) <= 5 {
		logger.Error(errors.New("Unable to find image Table."))
	}
	table := tableArray.Eq(4)
	rowsArray := table.Find("td")
	count := 0
	for rowIndex := range rowsArray.Nodes {
		var resultElement messageQueryBody
		row := rowsArray.Eq(rowIndex)
		textArray := row.Find("font")
		text := ""
		if len(textArray.Nodes) > 0 {
			text = textArray.Eq(0).Text()
		}
		imgArray := row.Find("img")
		if len(imgArray.Nodes) <= 0 {
			continue
		}
		imgTag := imgArray.Eq(0)
		imgLink, doLinkExist := imgTag.Attr("src")

		if !doLinkExist {
			continue
		}

		resultElement.Link = imgLink
		resultElement.Head = text
		resultObj = append(resultObj, resultElement)
		count = count + 1
		if count >= 10 {
			break
		}
	}
	return resultObj
}
