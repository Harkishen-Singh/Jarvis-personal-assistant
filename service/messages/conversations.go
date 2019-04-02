package messages

import (
	// "fmt"
	"net/http"
	// "encoding/json"
	// "os"
)

type response struct {
	username string
	message string
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
)

func loadJSONParsers() {

}

// GeneralConvHandler handles stuff related to general conversation
func GeneralConvHandler(req response, res http.ResponseWriter) {


}

