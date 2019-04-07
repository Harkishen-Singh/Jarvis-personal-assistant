package messages

import (
	"fmt"
	"net/http"
)

// HealthController controls tasks related to health services
func HealthController(medicine string,  res http.ResponseWriter) {

	fmt.Println("Health controllers")
}