package main

import (
	"fmt"
	"log"
	"net/http"
	"serveR/generate/controllers"
)

func main() {

	http.HandleFunc("/api/ping", controllers.PingHandler)
	http.HandleFunc("/api/login", controllers.HandleLogin)
	http.HandleFunc("/api/refresh", controllers.HandleRefresh)
	http.HandleFunc("/api/checktoken", controllers.HandleCheck)

	http.HandleFunc("/api/generate", controllers.GetGenerate)

	fmt.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
