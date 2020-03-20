package messages

import (
	"encoding/json"
	"fmt"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/logger"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type response struct {
	username string
	message  string
}

type jsonResponse struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Show    bool     `json:"show"`
	Result  []string `json:"result"`
}

// Messages json parser for default string types
type Messages struct {
	InitialGreetingsName  []string `json:"initial-greetings-name"`
	InitialGreetingsPlain []string `json:"initial-greetings-plain"`
	Help                  []string `json:"help"`
	About                 []string `json:"about"`
	Age                   []string `json:"age"`
	Birthday              []string `json:"birthday"`
}

// Messagesreplies json parser for default reply string types
type Messagesreplies struct {
	InitialGreetingsName  []string `json:"initial-greetings-name"`
	InitialGreetingsPlain []string `json:"initial-greetings-plain"`
	Help                  []string `json:"help"`
	About                 []string `json:"about"`
	Age                   []string `json:"age"`
	Birthday              []string `json:"birthday"`
}

var (
	messagesParser        Messages
	messagesRepliesParser Messagesreplies
	resp                  jsonResponse
	username, speak       string
	countMessage          int16
)

func init() {
	prefix := ""
	if os.Getenv("ENV") == "test" {
		prefix = "../"
	}

	fmt.Println("Loading messages JSON parsers....")
	messagesFile, err := os.Open(fmt.Sprintf("%sstatic/messages.json", prefix))
	messagesRepliesFile, err2 := os.Open(fmt.Sprintf("%sstatic/messages_replies.json", prefix))
	bytvalMF, _ := ioutil.ReadAll(messagesFile)
	bytvalMRF, _ := ioutil.ReadAll(messagesRepliesFile)
	if err != nil {
		logger.Error(err)
	}
	if err2 != nil {
		logger.Error(err2)
	}

	err1 := json.Unmarshal(bytvalMF, &messagesParser)
	err2 = json.Unmarshal(bytvalMRF, &messagesRepliesParser)
	if err1 != nil {
		logger.Error(err1)
	}
	if err2 != nil {
		logger.Error(err2)
	}

}

func filterForMessagesComparision(s string) (sr string) {

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
func GeneralConvHandler(req, name string, res http.ResponseWriter) string {

	fmt.Println("General conversation...")
	rand.Seed(time.Now().UnixNano())
	username = name
	message := filterForMessagesComparision(req)
	match := false

	// determine type of message
	if !match {
		isGreetingPlain := func(s string) bool {
			for i := 0; i < len(messagesParser.InitialGreetingsPlain); i++ {
				if strings.ToLower(s) == messagesParser.InitialGreetingsPlain[i] {
					match = true
					return true
				}
			}
			return false
		}(message)

		if isGreetingPlain {
			temp := greetingPlainController(message)
			resp = jsonResponse{true, temp, true, nil}
			speak = temp
			marshalled, _ := json.Marshal(resp)
			res.Write(marshalled)
		}
	}

	if !match {
		isGreetingName := func(s string) bool {
			for i := 0; i < len(messagesParser.InitialGreetingsName); i++ {
				if strings.ToLower(s) == messagesParser.InitialGreetingsName[i] {
					fmt.Println("contains ", strings.ToLower(s), " ", messagesParser.InitialGreetingsName[i])
					match = true
					return true
				}
			}
			return false
		}(message)

		if isGreetingName {
			temp := greetingNameController(message)
			resp = jsonResponse{true, temp, true, nil}
			speak = temp
			marshalled, _ := json.Marshal(resp)
			res.Write(marshalled)
		}
	}

	if !match {
		isHelp := func(s string) bool {
			for i := 0; i < len(messagesParser.Help); i++ {
				if strings.ToLower(s) == messagesParser.Help[i] {
					match = true
					return true
				}
			}
			return false
		}(message)

		if isHelp {
			temp := helpController(message)
			resp = jsonResponse{true, temp, true, nil}
			speak = temp
			marshalled, _ := json.Marshal(resp)
			res.Write(marshalled)
		}
	}

	if !match {
		isAbout := func(s string) bool {
			for i := 0; i < len(messagesParser.About); i++ {
				if strings.ToLower(s) == messagesParser.About[i] {
					match = true
					return true
				}
			}
			return false
		}(message)

		if isAbout {
			temp := aboutController(message)
			resp = jsonResponse{true, temp, true, nil}
			speak = temp
			marshalled, _ := json.Marshal(resp)
			res.Write(marshalled)
		}
	}

	if !match {
		fmt.Println("inside")
		isAge := func(s string) bool {
			for i := 0; i < len(messagesParser.Age); i++ {
				if strings.ToLower(s) == messagesParser.Age[i] {
					match = true
					return true
				}
			}
			return false
		}(message)

		if isAge {
			temp := ageController(message)
			fmt.Println("temp age is ", temp)
			resp = jsonResponse{true, temp, true, nil}
			speak = temp
			marshalled, _ := json.Marshal(resp)
			res.Write(marshalled)
		}
	}

	if !match {
		isBirthday := func(s string) bool {
			for i := 0; i < len(messagesParser.Birthday); i++ {
				if strings.ToLower(s) == messagesParser.Birthday[i] {
					match = true
					return true
				}
			}
			return false
		}(message)

		if isBirthday {
			temp := birthdayController(message)
			resp = jsonResponse{true, temp, true, nil}
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

func ageController(s string) string {

	numb := rand.Intn(len(messagesRepliesParser.Age))
	return messagesRepliesParser.Age[numb]

}

func birthdayController(s string) string {

	numb := rand.Intn(len(messagesRepliesParser.Birthday))
	return messagesRepliesParser.Birthday[numb]

}
