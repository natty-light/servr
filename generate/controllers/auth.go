package controllers

import (
	"net/http"
	"os"
	"serveR/generate/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

var user = map[string]string{
	os.Getenv("USER"): os.Getenv("PASS"),
}

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

	token, err := utils.CreateJWT(creds.Username)
	if err != nil {
		utils.AbortWithError(c, http.StatusInternalServerError, "Unable to produce signed JWT", err)
		return
	}

	c.JSON(http.StatusAccepted, token)
}

func HandleRefresh(c *gin.Context) {
	_, claims, err := utils.GetToken(c)
	if err != nil {
		utils.AbortWithError(c, http.StatusBadRequest, "Unable to retrieve token", err)
	}

	if time.Until(claims.ExpiresAt.Time) < 30*time.Second {
		utils.AbortWithError(c, http.StatusBadRequest, "Cannot refresh token with less than 30 seconds remaining", nil)
		return
	}

	token, err := utils.CreateJWT(claims.Username)
	if err != nil {
		utils.AbortWithError(c, http.StatusInternalServerError, "Unable to generate new JWT", err)
	}
	c.JSON(http.StatusAccepted, token)
}
