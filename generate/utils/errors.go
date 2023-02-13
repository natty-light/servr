package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ErrResponse struct {
	Err     error  `json:"error"`
	Message string `json:"message"`
}

func AbortWithError(c *gin.Context, code int, errMessage string, err error) {
	res := &ErrResponse{
		Err:     err,
		Message: errMessage,
	}
	fmt.Println(res)
	c.AbortWithStatusJSON(code, &res)
}
