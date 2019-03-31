package main

import (
	"github.com/hegedustibor/htgo-tts"
)


func main() {

	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	speech.Speak("Muskan is the best girl")
}