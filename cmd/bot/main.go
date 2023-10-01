package main

import (
	"github.com/filatkinen/tgbot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func main() {
	_ = godotenv.Load()

	apiKey := os.Getenv("BOTAPIKEY")
	if apiKey == "" {
		log.Printf("env BOTAPIKEY is not set")
		return
	}

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	commander := NewCommander(bot)

	for update := range updates {
		if update.Message != nil { // If we got a message

			switch update.Message.Command() {
			case "help":
				commander.HelpCommand(update.Message)
			case "list":
				commander.ListCommand(update.Message, productService)
			default:
				commander.DefaultBehavior(update.Message)
			}
		}
	}
}

type Commander struct {
	bot *tgbotapi.BotAPI
}

func NewCommander(bot *tgbotapi.BotAPI) *Commander {
	return &Commander{bot: bot}
}
func (c *Commander) HelpCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID,
		"/help - help\n"+
			"/list - list products")
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("got error while sending message: %s", err)
	}
}

func (c *Commander) ListCommand(message *tgbotapi.Message, productService *product.Service) {
	pr := strings.Builder{}
	pr.WriteString("Here all the products:\n\n")
	for _, v := range productService.List() {
		pr.WriteString(v.Title)
		pr.WriteString("\n")
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, pr.String())
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("got error while sending message: %s", err)
	}
}

func (c *Commander) DefaultBehavior(message *tgbotapi.Message) {
	//log.Printf("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, "You wrote:"+message.Text)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("got error while sending message: %s", err)
	}
}
