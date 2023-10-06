package lorry

import (
	"errors"
	"fmt"
	model "github.com/filatkinen/tgbot/internal/model/car/lorry"
	"github.com/filatkinen/tgbot/internal/service/car/lorry"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

const ListingCount = 3

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
	d := DummyLorryCommander{bot: bot, service: service}
	_, _ = d.service.Create(model.Lorry{Model: "One"})
	_, _ = d.service.Create(model.Lorry{Model: "Two"})
	_, _ = d.service.Create(model.Lorry{Model: "Three"})
	_, _ = d.service.Create(model.Lorry{Model: "Thor"})
	_, _ = d.service.Create(model.Lorry{Model: "Five"})
	_, _ = d.service.Create(model.Lorry{Model: "Six"})
	_, _ = d.service.Create(model.Lorry{Model: "Seven"})
	_, _ = d.service.Create(model.Lorry{Model: "Eight"})
	_, _ = d.service.Create(model.Lorry{Model: "Nine"})
	_, _ = d.service.Create(model.Lorry{Model: "Ten"})
	return &d
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
	if inputMsg.Message != nil {
		lorries, err := d.service.List(1, ListingCount)
		if err != nil {
			d.send(inputMsg.Message, err.Error())
			return
		}
		if len(lorries) == 0 {
			d.send(inputMsg.Message, "No records...")
		}
		sb := strings.Builder{}
		for _, l := range lorries {
			sb.WriteString(l.String())
			sb.WriteString("\n")
		}

		msg := tgbotapi.NewMessage(inputMsg.Message.Chat.ID, sb.String())
		kb := d.keyboardPrevNext("List", 1)
		msg.ReplyMarkup = kb
		_, err = d.bot.Send(msg)
		if err != nil {
			log.Printf("got error while sending message: %s", err)
		}
	} else if inputMsg.CallbackQuery != nil {
		fields := strings.Fields(inputMsg.CallbackQuery.Data)
		if len(fields) != 3 {
			return
		}
		idx, err := strconv.Atoi(fields[2])
		if err != nil {
			log.Printf("got error converting to int: %s", err)
			return
		}
		switch fields[1] {
		case "prev":
			idx -= ListingCount
		case "next":
			idx += ListingCount
		}

		lorries, err := d.service.List(uint64(idx), ListingCount)
		if err != nil && errors.Is(err, lorry.ErrWrongIndexSlice) {
			return
		}
		if len(lorries) == 0 {
			d.send(inputMsg.CallbackQuery.Message, "No records...")
		}
		sb := strings.Builder{}
		for _, l := range lorries {
			sb.WriteString(l.String())
			sb.WriteString("\n")
		}

		msg := tgbotapi.NewMessage(inputMsg.CallbackQuery.Message.Chat.ID, sb.String())
		kb := d.keyboardPrevNext("List", idx)
		msg.ReplyMarkup = kb
		_, err = d.bot.Send(msg)
		if err != nil {
			log.Printf("got error while sending message: %s", err)
		}
		return
	}
}

func (d DummyLorryCommander) keyboardPrevNext(prefix string, firstIdx int) tgbotapi.InlineKeyboardMarkup {
	data1 := prefix + " prev " + strconv.Itoa(firstIdx)
	data2 := prefix + " next " + strconv.Itoa(firstIdx)

	kb := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("<<", data1),
			tgbotapi.NewInlineKeyboardButtonData(">>", data2),
		),
	)
	return kb
}
func (d DummyLorryCommander) Delete(inputMsg *tgbotapi.Update) {
	if inputMsg.Message == nil {
		return
	}

	args := inputMsg.Message.CommandArguments()
	if args == "" {
		d.send(inputMsg.Message, "Wrong using of get command. Use: /delete_car_lorry ID")
		return
	}

	id, err := strconv.Atoi(args)
	if err != nil {
		d.send(inputMsg.Message, "Wrong number format of ID")
		return
	}
	_, err = d.service.Remove(uint64(id))
	if err != nil {
		d.send(inputMsg.Message, err.Error())
		return
	}
	d.send(inputMsg.Message, fmt.Sprintf("Deleted lorry with ID: %d", id))
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
	if inputMsg.Message == nil {
		return
	}
	args := inputMsg.Message.CommandArguments()
	argsSlice := strings.Fields(args)
	if args == "" || len(argsSlice) != 2 {
		d.send(inputMsg.Message, "Wrong using of get command. Use: /edit_car_lorry ID Name")
		return
	}

	id, err := strconv.Atoi(argsSlice[0])
	if err != nil {
		d.send(inputMsg.Message, "Wrong number format of ID")
		return
	}

	oldLorry, err := d.service.Describe(uint64(id))
	if err != nil {
		d.send(inputMsg.Message, err.Error())
		return
	}
	err = d.service.Update(uint64(id), model.Lorry{Model: argsSlice[1]})
	if err != nil {
		d.send(inputMsg.Message, err.Error())
		return
	}
	newLorry, err := d.service.Describe(uint64(id))
	if err != nil {
		d.send(inputMsg.Message, err.Error())
		return
	}

	d.send(inputMsg.Message, fmt.Sprintf("Success changing Old Lorry: %s, new Lorry:%s",
		oldLorry.String(), newLorry.String()))
}

func (d DummyLorryCommander) send(inputMsg *tgbotapi.Message, out string) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, out)
	_, err := d.bot.Send(msg)
	if err != nil {
		log.Printf("got error while sending message: %s", err)
	}
}
