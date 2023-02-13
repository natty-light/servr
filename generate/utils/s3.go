package utils

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func CreateS3Uploader() (*s3manager.Uploader, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION"))},
	)
	if err != nil {
		fmt.Println("ERROR: Unable to connect to AWS")
		return nil, err
	}
	uploader := s3manager.NewUploader(sess)

	return uploader, nil
}
