package messages

import (
	"fmt"
	"net/http"
	"encoding/json"
	"os"
	"io/ioutil"
	"math/rand"
	"strings"
)

type response struct {
	username string
	message string
}

type jsonResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Show bool `json:"show"`
	Result []string `json:"result"`
}

// Messages json parser for default string types
type Messages struct {
	InitialGreetingsName []string `json:"initial-greetings-name"`
	InitialGreetingsPlain []string `json:"initial-greetings-plain"`
	Help []string `json:"help"`
	About []string `json:"about"`
}

// Messagesreplies json parser for default reply string types
type Messagesreplies struct {
	InitialGreetingsName []string `json:"initial-greetings-name"`
	InitialGreetingsPlain []string `json:"initial-greetings-plain"`
	Help []string `json:"help"`
	About []string `json:"about"`
}

var (
	messagesParser Messages
	messagesRepliesParser Messagesreplies
	resp jsonResponse
	username, speak string
)

func loadJSONParsers(name string) {

	fmt.Println("Loading JSON parsers....")
	messagesFile, err := os.Open("messages/messages.json")
	messagesRepliesFile, err2 := os.Open("messages/messages_replies.json")
	bytvalMF, _ := ioutil.ReadAll(messagesFile)
	bytvalMRF, _ := ioutil.ReadAll(messagesRepliesFile)
	if err != nil   {
		panic(err)
	}
	if err2 != nil   {
		panic(err2)
	}
	username = name

	json.Unmarshal(bytvalMF, &messagesParser)
	json.Unmarshal(bytvalMRF, &messagesRepliesParser)

}

func filterForMessagesComparision(s string) (sr string)  {

	sr = strings.Replace(s, "?", " ", -1)
	sr = strings.Replace(sr, "%", " ", -1)
	sr = strings.Replace(sr, "#", " ", -1)
	sr = strings.Replace(sr, "$", " ", -1)
	sr = strings.Replace(sr, "@", " ", -1)
	sr = strings.Replace(sr, "&", " ", -1)
	sr = strings.Replace(sr, "^", " ", -1)
	sr = strings.Replace(sr, "*", " ", -1)
	return
}

// GeneralConvHandler handles stuff related to general conversation
func GeneralConvHandler(req, name string,  res http.ResponseWriter) string {

	fmt.Println("General conversation...")
	loadJSONParsers(name)
	message := filterForMessagesComparision(req)
	match := false

	// determine type of message
	if !match {
		isGreetingPlain := func(s string) bool {
			for i:=0; i< len(messagesParser.InitialGreetingsPlain); i++ {
				if strings.Contains(s, messagesParser.InitialGreetingsPlain[i]) {
					match = true
					return true
				}
			}
			return false
		}(message)

		if isGreetingPlain {
			temp := greetingPlainController(message)
			resp.Status = true
			resp.Show = true
			resp.Message = temp
			speak = temp
			marshalled, _ := json.Marshal(resp)
			res.Write(marshalled)
		}
	}


	if !match {
		isGreetingName := func(s string) bool {
			for i:=0; i< len(messagesParser.InitialGreetingsName); i++ {
				if strings.Contains(s, messagesParser.InitialGreetingsName[i]) {
					match = true
					return true
				}
			}
			return false
		}(message)

		if isGreetingName {
			temp := greetingNameController(message)
			resp.Status = true
			resp.Show = true
			resp.Message = temp
			speak = temp
			marshalled, _ := json.Marshal(resp)
			res.Write(marshalled)
		}
	}

	if !match {
		isHelp := func(s string) bool {
			for i:=0; i< len(messagesParser.Help); i++ {
				if strings.Contains(s, messagesParser.Help[i]) {
					match = true
					return true
				}
			}
			return false
		}(message)

		if isHelp {
			temp := helpController(message)
			resp.Status = true
			resp.Show = true
			resp.Message = temp
			speak = temp
			marshalled, _ := json.Marshal(resp)
			res.Write(marshalled)
		}
	}

	if !match {
		isAbout := func(s string) bool {
			for i:=0; i< len(messagesParser.About); i++ {
				if strings.Contains(s, messagesParser.About[i]) {
					match = true
					return true
				}
			}
			return false
		}(message)

		if isAbout {
			temp := aboutController(message)
			resp.Status = true
			resp.Show = true
			resp.Message = temp
			speak = temp
			marshalled, _ := json.Marshal(resp)
			res.Write(marshalled)
		}
	}

	return speak

}

func helpController(s string) string {

	numb := rand.Intn(len(messagesRepliesParser.Help))
	return messagesRepliesParser.Help[numb]

}

func greetingPlainController(s string) string {

	numb := rand.Intn(len(messagesRepliesParser.InitialGreetingsPlain))
	return messagesRepliesParser.InitialGreetingsPlain[numb]

}

func greetingNameController(s string) string {

	numb := rand.Intn(len(messagesRepliesParser.InitialGreetingsName))
	temp := messagesRepliesParser.InitialGreetingsName[numb]
	reply := fmt.Sprintf(temp, username) // note the formatter in messages_replies used
	return reply

}

func aboutController(s string) string {

	numb := rand.Intn(len(messagesRepliesParser.About))
	return messagesRepliesParser.About[numb]

}