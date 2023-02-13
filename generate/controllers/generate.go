package controllers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"serveR/generate/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GenerateRequestBody struct {
	Schools []string `json:"schools" binding:"required"`
}

func GetGenerate(c *gin.Context) {

	var request *GenerateRequestBody = &GenerateRequestBody{}
	if err := c.BindJSON(request); err != nil {
		utils.AbortWithError(c, http.StatusBadRequest, "Error: Unable to bind request", err)
		return
	}

	var id string = uuid.NewString()
	fileName, err := utils.Write(request.Schools, id)
	if err != nil {
		utils.AbortWithError(c, http.StatusInternalServerError, "Error: Unable to write to file system", err)
		return
	}

	// Set up channel to await R script completion
	ch := make(chan error)

	var script string = `./scripts/test.R`

	// Asynchronously run R Script
	go func() {
		cmd := exec.Command(`Rscript`, script, fileName)
		ch <- cmd.Run()
	}()

	// Initialize S3 Uploader
	uploader, err := utils.CreateS3Uploader()
	if err != nil {
		utils.AbortWithError(c, http.StatusInternalServerError, "Error: Unable to initialize S3 Uploader", err)
		return
	}

	// Create S3 UploadInput interface
	var input *s3manager.UploadInput = &s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
	}

	// Await completion of R Script using <-ch blocking feature
	err = <-ch
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error running R Script"+err.Error())
	} else {

		var outputFileName = id + ".pdf"

		// If we successfully run R Script, get the produced file
		file, err := os.Open(fileName)
		if err != nil {
			utils.AbortWithError(c, http.StatusInternalServerError, fmt.Sprintf("Error: Unable to open file %s \n", outputFileName), err)
			return
		}

		// create uploader
		input.Body, input.Key = file, aws.String((fileName))
		_, err = uploader.Upload(input)
		if err != nil {
			utils.AbortWithError(c, http.StatusInternalServerError, fmt.Sprintf("ERROR: Unable to upload file %s to S3 \n", outputFileName), err)
			return
		}

		c.JSON(http.StatusAccepted, outputFileName)
		os.Remove(outputFileName)
	}
	os.Remove(fileName)
}
