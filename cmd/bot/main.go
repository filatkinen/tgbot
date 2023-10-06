package main

import (
	commander "github.com/filatkinen/tgbot/internal/app/commands/car/lorry"
	"github.com/filatkinen/tgbot/internal/service/car/lorry"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

// Send any text message to the bot after the bot has been started

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

	productService := lorry.NewDummyLorryService()

	dummycommander := commander.NewLorryCommander(bot, productService)

	for update := range updates {
		HandlerMessage(bot, dummycommander, &update)
	}
}

func HandlerMessage(bot *tgbotapi.BotAPI, commander commander.LorryCommander, update *tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")

		if _, err := bot.Request(callback); err != nil {
			log.Printf("got error while sending callback message: %s", err)
		}
		switch {
		case strings.HasPrefix(update.CallbackQuery.Data, "List"):
			commander.List(update)
		}
	case update.Message != nil && update.Message.IsCommand():
		switch update.Message.Command() {
		case "list_car_lorry":
			commander.List(update)
		case "help_car_lorry":
			commander.Help(update)
		case "get_car_lorry":
			commander.Get(update)
		case "edit_car_lorry":
			commander.Edit(update)
		case "new_car_lorry":
			commander.New(update)
		case "delete_car_lorry":
			commander.Delete(update)
		}
	}
}
