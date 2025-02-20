package main

import (
	"backupManager/configs"
	"backupManager/handlers"
	object_storage "backupManager/object-storage"
	"backupManager/scheduler"
	"backupManager/telegram_bot"
	"backupManager/worker"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
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

	if configs.AppConfig.UploadToTelegram {
		go func() {
			RegisterTelegramEventDispatcher()
		}()
	}

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

func RegisterTelegramEventDispatcher() {
	bot := telegram_bot.GetTelegramBot()
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		telegram_bot.AddChatId(update.Message.Chat.ID)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Done")
		_, err := bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}

}
