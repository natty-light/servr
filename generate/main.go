package main

import (
	"fmt"
	"log"
	"net/http"
	"serveR/generate/controllers"
	"serveR/generate/utils"
)

func main() {

	pingHandler := http.HandlerFunc(controllers.PingHandler)
	loginHandler := http.HandlerFunc(controllers.HandleLogin)
	refreshHandler := http.HandlerFunc(controllers.HandleRefresh)
	checkHandler := http.HandlerFunc(controllers.HandleCheck)
	generateHandler := http.HandlerFunc(controllers.HandleGenerate)

	http.Handle("/api/ping", utils.EnableCors(pingHandler))
	http.Handle("/api/login", utils.EnableCors(loginHandler))
	http.Handle("/api/refresh", utils.EnableCors(refreshHandler))
	http.Handle("/api/checktoken", utils.EnableCors(checkHandler))
	http.Handle("/api/generate", utils.EnableCors(generateHandler))

	fmt.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
