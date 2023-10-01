package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func (c *Commander) HelpCommand(message *tgbotapi.Message) {
	helpMessage := strings.Builder{}
	for k, v := range registeredCommands {
		helpMessage.WriteString("/")
		helpMessage.WriteString(k)
		helpMessage.WriteString(" - ")
		helpMessage.WriteString(v.description)
		helpMessage.WriteString("\n")
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, helpMessage.String())
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("got error while sending message: %s", err)
	}
}

func init() {
	registeredCommands["help"] = command{f: (*Commander).HelpCommand,
		description: "show help info"}
}
