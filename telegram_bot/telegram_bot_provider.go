package telegram_bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var bot *tgbotapi.BotAPI

func GetTelegramBot() *tgbotapi.BotAPI {
	return bot
}

func init() {
	var err error
	bot, err = tgbotapi.NewBotAPI("7883182829:AAGEmPVev0aI0y5NPnR9xxPz1t2UTxCZwH0")
	if err != nil {
		panic(err)
	}

	bot.Debug = true
}
