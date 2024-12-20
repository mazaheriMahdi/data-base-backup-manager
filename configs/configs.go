package configs

import "os"

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
}

func init() {
	AppConfig.S3AccessKeyID = os.Getenv("S3_ACCESS_KEY_ID")
	AppConfig.S3SecretAccessKey = os.Getenv("S3_SECRET_ACCESS_KEY")
	AppConfig.S3Region = os.Getenv("S3_REGION")
	AppConfig.S3Bucket = os.Getenv("S3_BUCKET")
	AppConfig.S3Endpoint = os.Getenv("S3_ENDPOINT")

	AppConfig.BackUpDbUser = os.Getenv("BACKUP_DB_USER")
	AppConfig.BackUpDbPassword = os.Getenv("BACKUP_DB_PASSWORD")
	AppConfig.BackUpDbHost = os.Getenv("BACKUP_DB_HOST")
	AppConfig.BackUpDbPort = os.Getenv("BACKUP_DB_PORT")
	AppConfig.BackUpDbName = os.Getenv("BACKUP_DB_NAME")
}
