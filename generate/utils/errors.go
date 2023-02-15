package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrResponse struct {
	Err     error  `json:"error"`
	Message string `json:"message"`
}

func AbortWithError(w http.ResponseWriter, code int, errMessage string, err error) {
	res := &ErrResponse{
		Err:     err,
		Message: errMessage,
	}
	fmt.Println(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}

func AbortWithStatus(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}
