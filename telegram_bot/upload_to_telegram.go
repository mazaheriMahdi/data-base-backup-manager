package telegram_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"sync"
)

func UploadToTelegram(file *os.File) {
	backupFile := tgbotapi.FileReader{
		Name:   file.Name(),
		Reader: file,
	}
	bot := GetTelegramBot()
	var wait sync.WaitGroup
	for _, chatId := range GetAllChatIds() {
		wait.Add(1)
		go func() {
			backupMsg := tgbotapi.NewDocument(chatId, backupFile)
			_, err := bot.Send(backupMsg)
			if err != nil {
				log.Println(err)
			}
			wait.Done()
		}()
	}
	wait.Wait()
}
