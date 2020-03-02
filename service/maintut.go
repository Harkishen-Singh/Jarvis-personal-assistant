package main

import (
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/config"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/utils"
)

func main() {
	// utils.LoggerWarn("starting Jarvis service ...")
	// initiate service
	config.Init()
	utils.Server("3000")
}
