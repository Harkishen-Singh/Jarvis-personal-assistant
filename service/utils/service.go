package utils

import (
	"net/http"

	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/controllers"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/logger"
	// "../controllers"
)

func routes() {
	http.HandleFunc("/", controllers.HomeController)
	http.HandleFunc("/message", controllers.MessagesController)
	http.HandleFunc("/reminder", controllers.ReminderController)
	http.HandleFunc("/email", controllers.EmailController)
}

// Server service server for Jarvis
func Server(port string) {
	routes()
	// for logging the server status
	logger.Critic(http.ListenAndServe(":"+port, nil))
}
