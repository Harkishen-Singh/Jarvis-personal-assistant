package controllers

import (
	"fmt"
	"strings"
	"time"

	scrapper "github.com/Harkishen-Singh/Jarvis-personal-assistant/service/utils"
)

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
