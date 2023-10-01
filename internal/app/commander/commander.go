package commander

import (
	"github.com/filatkinen/tgbot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

func (c *Commander) HandlerMessage(message *tgbotapi.Message) {

	handler, ok := registeredCommands[message.Command()]
	if ok {
		handler.f(c, message)
	} else {
		c.DefaultBehavior(message)
	}
}
