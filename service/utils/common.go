package utils

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/config"
)

func randomUserAgent() string {
	userAgents := config.Get().UserAgents
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func getScrapeClient(proxyString interface{}) *http.Client {

	switch v := proxyString.(type) {
	case string:
		proxyUrl, _ := url.Parse(v)
		return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	default:
		return &http.Client{}

	}
}

func ScrapeClientRequest(searchURL string, proxyString interface{}) (*http.Response, error) {

	baseClient := getScrapeClient(proxyString)
	req, _ := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("User-Agent", randomUserAgent())

	res, err := baseClient.Do(req)
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("scraper received a non-200 status code suggesting a ban")
	}

	if err != nil {
		return nil, err
	}
	return res, nil
}
