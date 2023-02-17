package utils

import (
	"net/http"
	"os"
)

var allowedOrigin string = os.Getenv("ALLOWED_ORIGIN")

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", allowedOrigin)
}
