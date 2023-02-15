package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type DestructuredToken struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Expires int64  `json:"expiry"`
}

func CreateJWT(username string) (*DestructuredToken, error) {

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
	res := &DestructuredToken{Name: "token", Value: tokenString, Expires: expirationTime.UnixMilli()}
	return res, err
}

func CheckJWT(c *gin.Context) error {

	token, _, err := GetToken(c)
	if err != nil {
		return err
	}

	if token.Valid {
		return nil
	}
	return fmt.Errorf("unable to check JWT")
}

func GetToken(c *gin.Context) (*jwt.Token, *Claims, error) {
	tokenString := c.Query("token")
	if tokenString == "" {
		headerValue := strings.Split(c.Request.Header.Get("token"), " ")
		tokenString = headerValue[len(headerValue)-1]
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, nil, err
	}
	return token, claims, nil
}
