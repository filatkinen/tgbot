package main

import (
	botcommander "github.com/filatkinen/tgbot/internal/app/commander"
	"github.com/filatkinen/tgbot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
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

	updates := bot.GetUpdatesChan(tgbotapi.UpdateConfig{
		Offset:         0,
		Limit:          0,
		Timeout:        60,
		AllowedUpdates: nil,
	})

	productService := product.NewService()

	commander := botcommander.NewCommander(bot, productService)

	for update := range updates {
		commander.HandlerMessage(update)
	}
}
