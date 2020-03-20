package main

import (
	"os"

	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/utils"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	utils.Server(port)
}
