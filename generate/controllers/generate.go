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
	// var outputFileName = id + ".pdf"
	var outputFileName = "data.txt"
	fileName, err := utils.Write(request.Schools, id)

	if err != nil {
		utils.AbortWithError(c, http.StatusInternalServerError, "Error: Unable to write to file system", err)
		return
	}

	// Set up channel to await R script completion
	ch := make(chan error)

	var script string = fmt.Sprintf(`./scripts/%s`, os.Getenv("R_SCRIPT_NAME"))

	// Asynchronously run R Script
	go func() {
		cmd := exec.Command(`Rscript`, script, fileName)
		ch <- cmd.Run()
	}()

	// Initialize S3 Uploader
	uploader, err := utils.CreateS3Uploader()
	if err != nil {
		utils.AbortWithError(c, http.StatusInternalServerError, "Error: Unable to initialize S3 Uploader", err)
		os.Remove(fileName)
		return
	}

	// Create S3 UploadInput interface
	var bucket = os.Getenv("AWS_BUCKET_NAME")
	var input *s3manager.UploadInput = &s3manager.UploadInput{
		Bucket: aws.String(bucket),
	}

	// Await completion of R Script using <-ch blocking feature
	err = <-ch
	if err != nil {
		utils.AbortWithError(c, http.StatusInternalServerError, "Error: Failed to run R Script", err)
		os.Remove(fileName)
	} else {

		// If we successfully run R Script, get the produced file
		file, err := os.Open(outputFileName)
		if err != nil {
			utils.AbortWithError(c, http.StatusInternalServerError, fmt.Sprintf("Error: Unable to open file %s \n", outputFileName), err)
			os.Remove(outputFileName)
			os.Remove(fileName)
			return
		}

		// create uploader
		input.Body, input.Key = file, aws.String((outputFileName))
		_, err = uploader.Upload(input)
		if err != nil {
			utils.AbortWithError(c, http.StatusInternalServerError, fmt.Sprintf("ERROR: Unable to upload file %s to S3 \n", outputFileName), err)
			return
		}

		c.JSON(http.StatusAccepted, utils.GenerateS3ObjectURL(bucket, outputFileName))
	}
	os.Remove(outputFileName)
	os.Remove(fileName)
}
