package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"log"
)

func (c *Commander) DefaultBehavior(message *tgbotapi.Message) {
	//log.Printf("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, "You wrote: \""+message.Text+`"`)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("got error while sending message: %s", err)
	}
	log.Printf("\n-----\nreturn message: %+v\n-----", msg)
}
