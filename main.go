package main

import (
	"backupManager/configs"
	"backupManager/handlers"
	object_storage "backupManager/object-storage"
	"backupManager/scheduler"
	"backupManager/worker"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	session := object_storage.GenerateS3Session(configs.AppConfig.S3Region,
		configs.AppConfig.S3AccessKeyID,
		configs.AppConfig.S3SecretAccessKey,
		configs.AppConfig.S3Endpoint)

	workerInstance := worker.NewWorker()

	schedulerInstance := scheduler.NewScheduler()
	schedulerInstance.AddWorker(workerInstance)
	workerInstance.Start()
	schedulerInstance.Run()

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.POST("/Tasks", func(c *gin.Context) {
		handlers.AddTaskHandler(c, schedulerInstance, session)
	})
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
