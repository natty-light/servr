package controllers

import (
	"net/http"
	"os"
	"serveR/generate/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var user = map[string]string{
	os.Getenv("USER"): os.Getenv("PASS"),
}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func HandleLogin(c *gin.Context) {
	var creds *Credentials

	if err := c.BindJSON(&creds); err != nil {
		utils.AbortWithError(c, http.StatusBadRequest, "Error: Unable to log in", err)
		return
	}

	expectedPass, ok := user[creds.Username]

	if !ok || expectedPass != creds.Password {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	// you would like it to contain.
	var claims *Claims = &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		utils.AbortWithError(c, http.StatusInternalServerError, "Unable to produce signed JWT", err)
		return
	}

	res := struct {
		Name    string `json:"name"`
		Value   string `json:"value"`
		Expires int64  `json:"expiry"`
	}{}

	res.Name, res.Value, res.Expires = "token", tokenString, expirationTime.UnixMilli()
	c.JSON(http.StatusAccepted, res)
}
