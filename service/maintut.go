package main

import (
	"./utils"
)


func main() {
	utils.LoggerWarn("starting Jarvis service ...")

	// initiate service
	Server("3000")
}
