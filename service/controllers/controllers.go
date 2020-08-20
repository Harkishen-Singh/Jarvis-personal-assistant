package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/logger"
	scrapper "github.com/Harkishen-Singh/Jarvis-personal-assistant/service/utils"
	"github.com/PuerkitoBio/goquery"
)

// HomeController controls `/`.
func HomeController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey user! -- from Jarvis %s", r.URL.Path[1:])
}

// MessagesController controls `/message`.
// This involves general communication the users wants to have
// with the personal assistant.
func MessagesController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()

	request := response{
		username: r.FormValue("username"),
		message:  r.FormValue("message"),
	}

	AST(request, w)
}

func filterForSpeech(s string) string {

	s =   strings.Replace(s, "?", " ", -1)
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

func meanings(word string) []meaningStr {
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

func weather(city string, state string) weatherStr {
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

func scrapeYahoo(query string) []messageQueryBody {
	url := "https://in.search.yahoo.com/search?p=" + query
	res, err := scrapper.ScrapeClientRequest(url, nil)
	if err != nil {
		logger.Error(err)
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		logger.Error(err)
	}
	var resultObj []messageQueryBody
	listArray := doc.Find("div#web > ol")
	if len(listArray.Nodes) <= 0 {
		logger.Error(errors.New("Unable to find <ol>."))
	}
	list := listArray.Eq(0)
	elemArray := list.Find("li")
	for elemIndex := range elemArray.Nodes {
		elem := elemArray.Eq(elemIndex)
		titleContainerArray := elem.Find("div.compTitle")
		if len(titleContainerArray.Nodes) <= 0 {
			continue
		}
		titleContainer := titleContainerArray.Eq(0)
		titleArray := titleContainer.Find("h3")
		if len(titleArray.Nodes) <= 0 {
			continue
		}
		title := strings.TrimSpace(titleArray.Eq(0).Text())

		linkArray := titleContainer.Find("div")
		if len(linkArray.Nodes) <= 0 {
			continue
		}
		link := "https://" + strings.TrimSpace(linkArray.Eq(0).Text())

		descArray := elem.Find("div.compText")
		desc := ""
		if len(descArray.Nodes) > 0 {
			desc = strings.TrimSpace(descArray.Eq(0).Text())
		}
		var resultElement messageQueryBody
		resultElement.Link = link
		resultElement.Head = title
		resultElement.Desc = desc
		resultElement.DescLink = link
		resultObj = append(resultObj, resultElement)
	}
	return resultObj
}
