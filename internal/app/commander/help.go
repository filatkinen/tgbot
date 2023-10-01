package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) HelpCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID,
		"/help - help\n"+
			"/list - list products")
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("got error while sending message: %s", err)
	}
}
