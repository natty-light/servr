package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AbortWithError(c *gin.Context, code int, errMessage string, err error) {
	fmt.Println(errMessage)
	c.AbortWithError(code, err)
}
