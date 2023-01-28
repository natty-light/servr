package controllers

import (
	"os/exec"
	"serveR/utils"

	"github.com/gin-gonic/gin"
)

func GetGenerate(c *gin.Context) {
	ch := make(chan error)

	go func() {
		cmd := exec.Command(`python`, `C:\Users\jagal\OneDrive\Desktop\serveR\test.py`)
		ch <- cmd.Run()
	}()

	err := <-ch
	if err != nil {
		c.JSON(500, err)
	} else {
		res := utils.Read()
		c.JSON(200, res)
	}
}
