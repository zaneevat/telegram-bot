package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/zaneevat/telegram-bot/internal/app/commands"
	"github.com/zaneevat/telegram-bot/internal/service/product"
	"log"
	"os"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	productService := product.NewService()
	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			commander.Help(update.Message)
		case "list":
			commander.List(update.Message)
		default:
			commander.Default(update.Message)
		}
	}
	fmt.Println("hello")
}
