package controllers

import (
	"os/exec"
	"serveR/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GenerateRequestBody struct {
	Schools []string `json:"schools" binding:"required"`
}

func GetGenerate(c *gin.Context) {

	var request *GenerateRequestBody = &GenerateRequestBody{}
	c.BindJSON(&request)

	var id string = uuid.NewString()
	utils.Write(request.Schools, id)

	ch := make(chan error)

	go func() {
		cmd := exec.Command(`Rscript`, `C:\Users\jagal\OneDrive\Desktop\serveR\scripts\test.R`)
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
