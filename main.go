package main

import (
	"log"
	"scrum-estimator/internal/pkg/appconfig"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// 1. load config
	config, err := appconfig.LoadFromEnv()
	checkErr(err)

	// 2. create bot
	bot, err := tgbotapi.NewBotAPI(config.TelegramAPIToken)
	checkErr(err)

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// 3. initialize updates channel
	var updates tgbotapi.UpdatesChannel
	// only for GetUpdates - clear all webhooks
	_, err = bot.Request(tgbotapi.DeleteWebhookConfig{})
	checkErr(err)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates = bot.GetUpdatesChan(u)

	// 4. iterate each update received from telegram
	for update := range updates {
		log.Printf("----------------------- START OF PROCESSING -----------------------")
		if update.Message != nil {
			log.Printf("======================== MESSAGE DETECTED ========================")
		} else if update.CallbackQuery != nil {
			log.Printf("======================== CALLBACKQUERY DETECTED ========================")
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
