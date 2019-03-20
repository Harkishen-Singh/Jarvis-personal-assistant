package main

import (
	"net/http"
	"./controllers"
)

func routes() {
	http.HandleFunc("/", controllers.HomeController)
}

// Server service server for Jarvis
func Server(port string) {
	routes()
	http.ListenAndServe(":" + port, nil)
}