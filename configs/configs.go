package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	S3AccessKeyID     string
	S3SecretAccessKey string
	S3Region          string
	S3Bucket          string
	S3Endpoint        string
	BackUpDbUser      string
	BackUpDbPassword  string
	BackUpDbHost      string
	BackUpDbPort      string
	BackUpDbName      string
	UploadToS3        bool
	UploadToTelegram  bool
	TelegramBotToken  string
}

var AppConfig = Config{
	S3AccessKeyID:     "",
	S3SecretAccessKey: "",
	S3Region:          "",
	S3Bucket:          "",
	S3Endpoint:        "",
	BackUpDbUser:      "",
	BackUpDbPassword:  "",
	BackUpDbHost:      "",
	BackUpDbPort:      "",
	BackUpDbName:      "",
	UploadToS3:        false,
	UploadToTelegram:  false,
	TelegramBotToken:  "",
}

func init() {
	AppConfig.S3AccessKeyID = os.Getenv("S3_ACCESS_KEY_ID")
	AppConfig.S3SecretAccessKey = os.Getenv("S3_SECRET_ACCESS_KEY")
	AppConfig.S3Region = os.Getenv("S3_REGION")
	AppConfig.S3Bucket = os.Getenv("S3_BUCKET")
	AppConfig.S3Endpoint = os.Getenv("S3_ENDPOINT")
	var err error
	AppConfig.UploadToS3, err = strconv.ParseBool(os.Getenv("UPLOAD_TO_S3"))
	if err != nil {
		log.Fatalln(fmt.Errorf("faild to parse UPLOAD_TO_S3: %v", err))
	}
	AppConfig.UploadToTelegram, err = strconv.ParseBool(os.Getenv("UPLOAD_TO_TELEGRAM"))
	if err != nil {
		log.Fatalln(fmt.Errorf("faild to parse UPLOAD_TO_TELEGRAM: %v", err))
	}
	AppConfig.TelegramBotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
}
