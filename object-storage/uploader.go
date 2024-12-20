package object_storage

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"os"
)

func uploadFile(file *os.File, session *s3.S3) error {
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		return err
	}

	// TODO : print or put id of generated history to db
	_, err := session.PutObject(&s3.PutObjectInput{
		Body: bytes.NewReader(buf.Bytes()),
		Key:  aws.String(file.Name()),
	})
	if err != nil {
		return err
	}
	return nil
}
