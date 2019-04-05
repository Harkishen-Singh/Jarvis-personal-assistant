package main

import (
	"github.com/hegedustibor/htgo-tts"
)


func main() {

	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	speech.Speak("I can help you with web search, specific google yahoo bing search, images, videos and also find answers for any of your queries")
}