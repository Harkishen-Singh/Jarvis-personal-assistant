package utils

import (
	"github.com/hegedustibor/htgo-tts"
)

// TextToSpeech converts text to speech on a string, return status as boolean
func TextToSpeech(message string) bool {

    speech := htgotts.Speech{Folder: "audio", Language: "en"}
	speech.Speak(message)
	return true

}