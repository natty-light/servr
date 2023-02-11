package controllers

import (
	"net/http"
	"os"
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
	if err := c.BindJSON(request); err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}
	var id string = uuid.NewString()
	fileName, err := utils.Write(request.Schools, id)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ch := make(chan error)

	var script string = `./scripts/test.R`

	go func() {
		cmd := exec.Command(`Rscript`, script, fileName)
		ch <- cmd.Run()
	}()

	// Configure AWS S3 Here

	err = <-ch
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error running R Script"+err.Error())
	} else {
		res := utils.Read()
		c.JSON(http.StatusAccepted, res)
	}
	os.Remove(fileName)
}
