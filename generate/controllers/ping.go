package controllers

import (
	"encoding/json"
	"net/http"
	"serveR/generate/utils"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	body := struct{ Message string }{Message: "pong"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(body)
}
