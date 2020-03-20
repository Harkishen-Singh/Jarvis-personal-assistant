package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/logger"
)

// Configuration to contain all configuration
type Configuration struct {
	KeywordPriority []string `json:"keywordPriority"`
	UserAgents      []string `json:"userAgents"`
}

var config Configuration

// init reads and stores data from config.json
func init() {
	jsonFile, err := os.Open("../static/config.json")
	if err != nil {
		fmt.Print("I am here")
		logger.Error(err)
	}
	byteValue, ioError := ioutil.ReadAll(jsonFile)
	if ioError != nil {
		logger.Error(ioError)
	}
	json.Unmarshal(byteValue, &config)
	fmt.Println("loaded configuration")
	defer jsonFile.Close()
}

// Get returns the configuration
func Get() Configuration {
	return config
}
