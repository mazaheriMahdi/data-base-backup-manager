package object_storage

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GenerateS3Session(region, accessKey, secret, endpoint string) *s3.S3 {
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(region),
		Endpoint:         aws.String(endpoint),
		Credentials:      credentials.NewStaticCredentials(accessKey, secret, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return nil
	}
	return s3.New(sess)
}
