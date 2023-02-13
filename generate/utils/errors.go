package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AbortWithError(c *gin.Context, code int, errMessage string, err error) {
	fmt.Println(errMessage)
	fmt.Println(err.Error())
	c.AbortWithError(code, err)
}
