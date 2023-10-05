package commander

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) ListCommand(message *tgbotapi.Message) {
	pr := strings.Builder{}
	pr.WriteString("Here all the products:\n\n")
	for _, v := range c.productService.List() {
		pr.WriteString(v.Title)
		pr.WriteString("\n")
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, pr.String())

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Next Page", "some data"),
	))
	m, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("got error while sending message: %s", err)
	}
	log.Printf("\n-----\nreturn message: %+v\n-----", m)
}

func init() {
	registeredCommands["list"] = command{f: (*Commander).ListCommand,
		description: "get products list"}
}
