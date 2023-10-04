package lorry

import (
	"github.com/filatkinen/tgbot/internal/service/car/lorry"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type LorryCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type DummyLorryCommander struct {
	bot     *tgbotapi.BotAPI
	service lorry.LorryService
}

func NewLorryCommander(bot *tgbotapi.BotAPI, service lorry.LorryService) LorryCommander {
	return DummyLorryCommander{bot: bot, service: service}
}
