package lorry

import (
	model "github.com/filatkinen/tgbot/internal/model/car/lorry"
	"github.com/filatkinen/tgbot/internal/service/car/lorry"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

type LorryCommander interface {
	Help(inputMsg *tgbotapi.Update)
	Get(inputMsg *tgbotapi.Update)
	List(inputMsg *tgbotapi.Update)
	Delete(inputMsg *tgbotapi.Update)

	New(inputMsg *tgbotapi.Update)  // return error not implemented
	Edit(inputMsg *tgbotapi.Update) // return error not implemented
}

type DummyLorryCommander struct {
	bot     *tgbotapi.BotAPI
	service lorry.LorryService
}

func NewLorryCommander(bot *tgbotapi.BotAPI, service lorry.LorryService) LorryCommander {
	return DummyLorryCommander{bot: bot, service: service}
}

func (d DummyLorryCommander) Help(inputMsg *tgbotapi.Update) {
	out := strings.Builder{}
	out.WriteString("/help_car_lorry - show help\n")
	out.WriteString("/get_car_lorry ID - show lorry by ID\n")
	out.WriteString("/list_car_lorry  list lorries\n")
	out.WriteString("/delete_car_lorry ID -  delete lorry by ID\n")
	out.WriteString("/new_car_lorry Name - adding new lorry with Name\n")
	out.WriteString("/edit_car_lorry ID Name - editing lorry: setting new Name by ID\n")

	msg := tgbotapi.NewMessage(inputMsg.Message.Chat.ID, out.String())
	_, err := d.bot.Send(msg)
	if err != nil {
		log.Printf("got error while sending message from Help function: %s", err)
	}
}

func (d DummyLorryCommander) Get(inputMsg *tgbotapi.Update) {
	if inputMsg.Message == nil {
		return
	}
	args := inputMsg.Message.CommandArguments()
	if args == "" {
		d.send(inputMsg.Message, "Wrong using of get command. Use: /get_car_lorry ID")
		return
	}
	id, err := strconv.Atoi(args)
	if err != nil {
		d.send(inputMsg.Message, "Wrong number format of ID")
		return
	}
	l, err := d.service.Describe(uint64(id))
	if err != nil {
		d.send(inputMsg.Message, err.Error())
		return
	}
	d.send(inputMsg.Message, l.String())
}

func (d DummyLorryCommander) List(inputMsg *tgbotapi.Update) {
	//TODO implement me
	panic("implement me")
}

func (d DummyLorryCommander) Delete(inputMsg *tgbotapi.Update) {
	//TODO implement me
	panic("implement me")
}

func (d DummyLorryCommander) New(inputMsg *tgbotapi.Update) {
	if inputMsg.Message == nil {
		return
	}
	args := inputMsg.Message.CommandArguments()
	if args == "" {
		d.send(inputMsg.Message, "Wrong using of get command. Use: /new_car_lorry Name")
		return
	}
	id, err := d.service.Create(model.Lorry{Model: args})
	if err != nil {
		d.send(inputMsg.Message, err.Error())
		return
	}
	l, err := d.service.Describe(id)
	if err != nil {
		d.send(inputMsg.Message, err.Error())
		return
	}
	d.send(inputMsg.Message, "Created new lorry:"+l.String())
}

func (d DummyLorryCommander) Edit(inputMsg *tgbotapi.Update) {
	//TODO implement me
	panic("implement me")
}

func (d DummyLorryCommander) send(inputMsg *tgbotapi.Message, out string) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, out)
	_, err := d.bot.Send(msg)
	if err != nil {
		log.Printf("got error while sending message: %s", err)
	}
}
