package messages

import (
	"fmt"
	"net/http"
	"encoding/json"
	"os"
	"io/ioutil"
	"math/rand"
)

type response struct {
	username string
	message string
}

type jsonResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Result []string `json:"result"`
}

// Messages json parser for default string types
type Messages struct {
	InitialGreetingsName []string `json:"initial-greetings-name"`
	InitialGreetingsPlain []string `json:"initial-greetings-plain"`
}

// Messagesreplies json parser for default reply string types
type Messagesreplies struct {
	InitialGreetingsName []string `json:"initial-greetings-name"`
	InitialGreetingsPlain []string `json:"initial-greetings-plain"`
}

var (
	messagesParser Messages
	messagesRepliesParser Messagesreplies
	resp jsonResponse
	username string
)

func loadJSONParsers(name string) {

	fmt.Println("Loading JSON parsers....")
	messagesFile, err := os.Open("messages.json")
	messagesRepliesFile, err2 := os.Open("messages_replies.json")
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

// GeneralConvHandler handles stuff related to general conversation
func GeneralConvHandler(req response, res http.ResponseWriter) {

	loadJSONParsers(req.username)
	message := req.message
	// determine type of message
	isGreetingPlain := func(s string) bool {
		for i:=0; i< len(messagesParser.InitialGreetingsPlain); i++ {
			if messagesParser.InitialGreetingsPlain[i] == s {
				return true
			}
		}
		return false
	}(message)

	if isGreetingPlain {
		temp := greetingPlainController(message)
		resp.Status = true
		resp.Message = temp
		marshalled, _ := json.Marshal(resp)
		res.Write(marshalled)
		return ;
	}

	isGreetingName := func(s string) bool {
		for i:=0; i< len(messagesParser.InitialGreetingsName); i++ {
			if messagesParser.InitialGreetingsName[i] == s {
				return true
			}
		}
		return false
	}(message)

	if isGreetingName {
		temp := greetingNameController(message)
		resp.Status = true
		resp.Message = temp
		marshalled, _ := json.Marshal(resp)
		res.Write(marshalled)
	}

}

func greetingPlainController(s string) string {

	numb := rand.Intn(len(messagesRepliesParser.InitialGreetingsPlain))
	return messagesRepliesParser.InitialGreetingsPlain[numb]

}

func greetingNameController(s string) string {

	numb := rand.Intn(len(messagesRepliesParser.InitialGreetingsName))
	reply := fmt.Sprintf(messagesRepliesParser.InitialGreetingsName[numb], &username) // note the formatter in messages_replies used
	return reply

}
