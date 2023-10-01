package commander

import (
	"github.com/filatkinen/tgbot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type command struct {
	f           func(*Commander, *tgbotapi.Message)
	description string
}

var registeredCommands = map[string]command{}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, service *product.Service) *Commander {
	return &Commander{bot: bot, productService: service}
}

func (c *Commander) HandlerMessage(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Data: "+update.CallbackQuery.Data)
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("got error while sending message: %s", err)
		}
		return
	}
	if update.Message == nil {
		return
	}
	com := update.Message.Command()
	handler, ok := registeredCommands[com]
	if ok {
		handler.f(c, update.Message)
	} else {
		c.DefaultBehavior(update.Message)
	}
}
