package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"os"
)

func main() {

	accessKey := "access-key"
	secretKey := "secret-key"

	cfg := aws.NewConfig().
		WithRegion("us-west-1").
		WithEndpoint("http://mountebank-host:mountebank-port").
		WithCredentials(credentials.NewStaticCredentials(accessKey, secretKey, ""))
	sess := session.Must(session.NewSession(cfg))

	// Create a new session with the credentials.
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(endpoints.ApSoutheast1RegionID), // Replace with the desired AWS region
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		log.Fatalf("failed to create new session %v", err.Error())
	}
	s3Client := s3.New(sess)

	bucketName := "qa-mf-kyc-images-bucket"
	filePath := "./test-image.png"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("idcard/e-aon.png"),
		Body:   file,
	})

	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}

	fmt.Println("Image uploaded successfully!")
}
