package telegram_bot

import "log"

var chatIds = []int64{}

func AddChatId(chatId int64) {
	log.Printf("Telegram Backup Bucket with id %s Added To Buckets \n", chatId)
	chatIds = append(chatIds, chatId)
}
func GetAllChatIds() []int64 {
	return chatIds
}
