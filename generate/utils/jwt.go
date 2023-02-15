package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

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

func CheckJWT(r *http.Request) error {

	token, _, err := GetToken(r)
	if err != nil {
		return err
	}

	if token.Valid {
		return nil
	}
	return fmt.Errorf("token invalid")
}

func GetToken(r *http.Request) (*jwt.Token, *Claims, error) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		return nil, nil, fmt.Errorf("auth header in wrong format")
	}

	tokenString := strings.TrimSpace(splitToken[1])

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
