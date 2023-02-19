package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"serveR/generate/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

type GenerateRequestBody struct {
	Schools []string `json:"schools"`
}

type GenerateResponseBody struct {
	Url string `json:"url"`
}

func HandleGenerate(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	switch r.Method {
	case "OPTIONS":
		fmt.Println("Handling Preflight Check")
		w.WriteHeader(http.StatusOK)
		return
	case "POST":
		fmt.Println("Handling Post Request")
		if err := utils.CheckJWT(r); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		request := GenerateRequestBody{}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			utils.AbortWithError(w, http.StatusBadRequest, "Error: Unable to bind request", err)
			return
		}

		var id string = uuid.NewString()
		var outputFileName = id + ".pdf"
		fileName, err := utils.Write(request.Schools, id)

		if err != nil {
			utils.AbortWithError(w, http.StatusInternalServerError, "Error: Unable to write to file system", err)
			return
		}

		// Set up channel to await R script completion
		ch := make(chan error)
		outch := make(chan []byte)

		var script string = fmt.Sprintf(`./scripts/%s`, os.Getenv("R_SCRIPT_NAME"))

		// Asynchronously run R Script
		go func() {
			cmd := exec.Command(`Rscript`, script, fileName)
			stdout, stderr := cmd.CombinedOutput()
			ch <- stderr
			outch <- stdout
			close(ch)
			close(outch)
		}()

		// Initialize S3 Uploader
		uploader, err := utils.CreateS3Uploader()
		if err != nil {
			utils.AbortWithError(w, http.StatusInternalServerError, "Error: Unable to initialize S3 Uploader", err)
			os.Remove(fileName)
			return
		}

		// Create S3 UploadInput interface
		var bucket = os.Getenv("AWS_BUCKET_NAME")
		var input *s3manager.UploadInput = &s3manager.UploadInput{
			Bucket:      aws.String(bucket),
			ContentType: aws.String("application/pdf"),
		}

		// Await completion of R Script using <-ch blocking feature
		err = <-ch
		out := <-outch
		if err != nil {
			fmt.Println(string(out))
			utils.AbortWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error: Failed to run R Script with message %s", string(out)), err)
			os.Remove(fileName)
		} else {
			// If we successfully run R Script, get the produced file
			file, err := os.Open(outputFileName)
			if err != nil {
				utils.AbortWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error: Unable to open file %s \n", outputFileName), err)
				os.Remove(outputFileName)
				os.Remove(fileName)
				return
			}

			// create uploader
			input.Body, input.Key = file, aws.String((outputFileName))
			_, err = uploader.Upload(input)
			if err != nil {
				utils.AbortWithError(w, http.StatusInternalServerError, fmt.Sprintf("ERROR: Unable to upload file %s to S3 \n", outputFileName), err)
				return
			}

			res := GenerateResponseBody{Url: utils.GenerateS3ObjectURL(*input.Bucket, outputFileName)}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(res)
		}
		os.Remove(outputFileName)
		os.Remove(fileName)

	}
}
