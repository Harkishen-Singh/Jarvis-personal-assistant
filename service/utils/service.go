package utils

import (
	"net/http"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/controllers"
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
	http.ListenAndServe(":" + port, nil)
}