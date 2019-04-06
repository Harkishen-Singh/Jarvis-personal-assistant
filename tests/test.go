package main

import (
	"fmt"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/controllers"
)


func main() {
	b := []string{"gooogle","googl" , "googleee", "goooogggle", "goole"}
	result := controllers.FindBestMatch("google", b)
	fmt.Println("Result: ", result)
}