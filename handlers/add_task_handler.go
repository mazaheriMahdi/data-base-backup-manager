package handlers

import (
	"backupManager/backup"
	"backupManager/configs"
	"backupManager/dump"
	"backupManager/scheduler"
	"backupManager/worker"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type AddTaskDto struct {
	Name   string `form:"name"`
	Offset int    `form:"offset"`
	dump.DBConfiguration
}

func getBackUp(session *s3.S3, configuration dump.DBConfiguration) error {
	err := backup.Run(configuration, session, configs.AppConfig.S3Bucket)
	if err != nil {
		return err
	}
	return nil
}
func AddTaskHandler(context *gin.Context, schedulerInstance *scheduler.Scheduler, session *s3.S3) {
	var addTaskDto AddTaskDto
	err := context.ShouldBindBodyWith(&addTaskDto, binding.JSON)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if addTaskDto.Name == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}
	if addTaskDto.Offset == 0 || addTaskDto.Offset < 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "offset should be positive and must be greater than 0"})
		return
	}
	if addTaskDto.DB == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "db name is required"})
		return
	}

	schedulerInstance.AddTask(&scheduler.ScheduledTask{
		Id:   uuid.New(),
		Name: addTaskDto.Name,
		Action: &worker.Task{
			Id: uuid.New(),
			Action: func() error {
				return getBackUp(session, addTaskDto.DBConfiguration)
			},
		},
		NextRun:   time.Time{},
		RunOffset: time.Duration(addTaskDto.Offset) * time.Second,
	})

}
