package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Configuration : Struct to contain all configuration
type Configuration struct {
	KeywordPriority []string `json:"keywordPriority"`
	UserAgents []string `json:"userAgents"`
}

var config Configuration

// Init reads and stores data from config.json
func Init() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	byteValue, ioError := ioutil.ReadAll(jsonFile)
	if ioError != nil {
		panic(ioError)
	}
	json.Unmarshal(byteValue, &config)
	fmt.Println("[JARVIS] config.json read successfully")
	defer jsonFile.Close()
}

// Get : It returns the configuration
func Get() Configuration {
	return config
}
