package main

import (
	"buy_list/commands"
	"buy_list/statuses"
	"buy_list/storage/postgresql"
	"buy_list/tgbot"
	timer "buy_list/timer"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}
	db := postgresql.Connect(os.Getenv("POSTGRESQL_TOKEN"))

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 90
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		user_nickname := update.SentFrom().UserName
		user_msg := update.Message.Text
		user_name := update.Message.From.FirstName
		user_id := update.Message.From.ID
		chat_id := update.Message.Chat.ID
		status := db.GetStatus(user_id, chat_id)

		log.Printf("%s (%s): %s \n", user_name, user_nickname, user_msg)
		if update.Message.IsCommand() {
			msg := commands.CommandHandler(user_msg, user_id, chat_id, user_nickname, user_name, db)
			tgbot.SendMessageToUser(bot, update, msg)
		} else {
			msg := statuses.StatusHandler(status, user_msg, user_id, chat_id, db)
			tgbot.SendMessageToUser(bot, update, msg)
		}
		db.UpdateTimer(user_id, chat_id)
		timer.SetTimerList(user_id, chat_id, bot, db, update)
	}

}
