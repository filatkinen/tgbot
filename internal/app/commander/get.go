package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Get(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "TBD")
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("got error while sending message: %s", err)
	}
}

func init() {
	registeredCommands["get"] = command{f: (*Commander).Get,
		description: "get (TBD)"}
}
