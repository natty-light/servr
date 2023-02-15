package controllers

import (
	"fmt"
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

type ResponseBody struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Expires int64  `json:"expiry"`
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

	res, err := IssueJWT(c, creds.Username)
	if err != nil {
		utils.AbortWithError(c, http.StatusInternalServerError, "Unable to produce signed JWT", err)
		return
	}

	c.JSON(http.StatusAccepted, res)
}

func HandleRefresh(c *gin.Context) {

}

func IssueJWT(c *gin.Context, username string) (*ResponseBody, error) {

	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, which includes the username and expiry time
	// you would like it to contain.
	var claims *Claims = &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	var res *ResponseBody = &ResponseBody{Name: "token", Value: tokenString, Expires: expirationTime.UnixMilli()}
	return res, nil
}

func CheckJWT(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtKey, nil
	})

	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	}

	return false, fmt.Errorf("unable to process token")
}
