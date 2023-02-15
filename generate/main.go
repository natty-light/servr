package main

import (
	"log"
	"net/http"
	"serveR/generate/controllers"
)

func main() {

	http.HandleFunc("/api/ping", controllers.PingHandler)
	http.HandleFunc("/api/login", controllers.HandleLogin)
	http.HandleFunc("/api/refresh", controllers.HandleRefresh)
	http.HandleFunc("/api/generate", controllers.GetGenerate)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
