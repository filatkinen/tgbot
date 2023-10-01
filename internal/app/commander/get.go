package commander

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(message *tgbotapi.Message) {
	send := func(s string) {
		msg := tgbotapi.NewMessage(message.Chat.ID, s)
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("got error while sending message: %s", err)
		}
	}

	args := message.CommandArguments()

	num, err := strconv.Atoi(args)
	if err != nil {
		send(fmt.Sprintf("wrong number format: %s", args))
		return
	}

	productTitle, err := c.productService.GetProductTitle(num)
	if err != nil {
		send(err.Error())
		return
	} else {
		send(productTitle)
		return
	}
}

func init() {
	registeredCommands["get"] = command{f: (*Commander).Get,
		description: "get (TBD)"}
}
