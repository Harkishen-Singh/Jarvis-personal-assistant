package main

import (
	"./utils"
)


func main() {
	utils.LoggerWarn("starting Jarvis service ...")

	// initiate service
	utils.Server("3000")
}
