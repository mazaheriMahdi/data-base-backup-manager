package telegram_bot

import (
	"backupManager/configs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func GetTelegramBot() *tgbotapi.BotAPI {
	return bot
}

func init() {
	var err error
	bot, err = tgbotapi.NewBotAPI(configs.AppConfig.TelegramBotToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true
}
