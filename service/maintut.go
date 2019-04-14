package main

import (
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/utils"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/controllers"
	"fmt"
)
	
func main() {
	// utils.LoggerWarn("starting Jarvis service ...")
	// initiate service

	result := controllers.CheckMail()
	if len(result) > 0 {
		go controllers.CheckTime(1, result)
	} else {
		fmt.Println("bye")
	}

	utils.Server("3000")
	fmt.Println("reached")
}
