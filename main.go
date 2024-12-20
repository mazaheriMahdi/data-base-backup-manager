package main

import (
	"backupManager/backup"
	"backupManager/configs"
	"backupManager/dump"
	object_storage "backupManager/object-storage"
	"backupManager/scheduler"
	"backupManager/worker"
	"github.com/google/uuid"
	"strconv"
	"time"
)

func main() {

	session := object_storage.GenerateS3Session(configs.AppConfig.S3Region,
		configs.AppConfig.S3AccessKeyID,
		configs.AppConfig.S3SecretAccessKey,
		configs.AppConfig.S3Endpoint)

	dbBackUpFunc := func() error {
		port, _ := strconv.Atoi(configs.AppConfig.BackUpDbPort)
		err := backup.Run(dump.DBConfiguration{
			Host:     configs.AppConfig.BackUpDbHost,
			Port:     port,
			DB:       configs.AppConfig.BackUpDbName,
			Username: configs.AppConfig.BackUpDbUser,
			Password: configs.AppConfig.BackUpDbPassword,
		}, session, configs.AppConfig.S3Bucket)
		if err != nil {
			return err
		}
		return nil
	}

	workerInstance := worker.NewWorker()
	workerInstance.Start()

	schedulerInstance := scheduler.NewScheduler()
	schedulerInstance.AddWorker(workerInstance)
	schedulerInstance.AddTask(&scheduler.ScheduledTask{
		Action: worker.Task{
			Id:     uuid.New(),
			Action: dbBackUpFunc,
		},
		NextRun:   time.Now(),
		RunOffset: 12 * time.Hour,
		Id:        uuid.New(),
		Name:      "Backup cms_production",
	})
	schedulerInstance.Run()

	for {
		time.Sleep(1 * time.Second)
	}
}
