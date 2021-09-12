package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	creds := credentials.NewStaticCredentials("AWS_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY", "")

	sess, _ := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String("ap-northeast-1")},
	)

	bucket := "test-common"
	filename := "fuga.txt"
	file, _ := os.Open(filename)
	defer file.Close()

	uploader := s3manager.NewUploader(sess)
	uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
}
