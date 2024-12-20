package backup

import (
	"backupManager/dump"
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"log"
	"os"
)

func Run(configuration dump.DBConfiguration, s3Session *s3.S3, bucket string) error {
	dumpFile, err := dump.GetDump(configuration)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, dumpFile); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		return err
	}
	log.Println("Upload Started....")
	_, err = s3Session.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fmt.Sprintf("%s-%s", configuration.DB, dumpFile.Name())),
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		return err
	}
	log.Println("Upload Finished....")
	err = os.Remove(dumpFile.Name())
	if err != nil {
		return err
	}
	return nil
}