package middleware

import (
	"net/http"
	"serveR/generate/utils"

	"github.com/gin-gonic/gin"
)

func AuthorizeToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := utils.CheckJWT(c); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		}
	}
}
