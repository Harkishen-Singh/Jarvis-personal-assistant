package controllers

import (
	"fmt"
	"github.com/hegedustibor/htgo-tts"
	"time"
)

// TextToSpeech converts text to speech on a string, return status as boolean
func TextToSpeech(message string, waitInSeconds int32) bool {

	time.Sleep(time.Duration(waitInSeconds))
	fmt.Println("Speaking -> ", message)
	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	speech.Speak(message)
	return true

}
