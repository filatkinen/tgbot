package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
)

func main() {
	_ = godotenv.Load()

	apiKey := os.Getenv("BOTAPIKEY")
	if apiKey == "" {
		log.Printf("env BOTAPIKEY is not set")
		return
	}

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(-100)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	//time.Sleep(time.Millisecond * 500)
	//for len(updates) != 0 {
	//	<-updates
	//}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		fmt.Printf("%+v", update.Message)
	}
}
