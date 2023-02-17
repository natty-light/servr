package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"serveR/generate/utils"
	"time"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

var user = map[string]string{
	os.Getenv("USER"): os.Getenv("PASS"),
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	var creds Credentials

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		utils.AbortWithError(w, http.StatusBadRequest, "Error: Unable to log in", err)
		return
	}

	expectedPass, ok := user[creds.Username]

	if !ok || expectedPass != creds.Password {
		utils.AbortWithStatus(w, http.StatusUnauthorized)
		return
	}

	token, err := utils.CreateJWT(creds.Username)
	if err != nil {
		utils.AbortWithError(w, http.StatusInternalServerError, "Unable to produce signed JWT", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(token)
}

func HandleRefresh(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	_, claims, err := utils.GetToken(r)
	if err != nil {
		utils.AbortWithError(w, http.StatusBadRequest, "Unable to retrieve token", err)
	}

	if time.Until(claims.ExpiresAt.Time) < 30*time.Second {
		utils.AbortWithError(w, http.StatusBadRequest, "Cannot refresh token with less than 30 seconds remaining", nil)
		return
	}

	token, err := utils.CreateJWT(claims.Username)
	if err != nil {
		utils.AbortWithError(w, http.StatusInternalServerError, "Unable to generate new JWT", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(token)
}

func HandleCheck(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	if err := utils.CheckJWT(r); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
