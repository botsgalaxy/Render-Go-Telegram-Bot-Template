package main

import (
	"log"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/botsgalaxy/render-go-telegram-bot-template/internal/database"
)

func init() {
	go database.InitDB()
}
func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("TOKEN environment variable is empty")
	}

	appUrl := os.Getenv("APP_URL")
	if appUrl == "" {
		panic("APP_URL environment variable is empty")
	}

	webhookSecret := os.Getenv("WEBHOOK_SECRET")
	if webhookSecret == "" {
		panic("WEBHOOK_SECRET environment variable is empty")
	}

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT environment variable is empty")

	}

	b, err := gotgbot.NewBot(token, nil)
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)

	addHandlers(dispatcher)

	webhookOpts := ext.WebhookOpts{
		ListenAddr:  "0.0.0.0:" + port,
		SecretToken: webhookSecret,
	}

	// The bot's urlPath can be anything. Here, we use "custom-path/<TOKEN>" as an example.
	// It can be a good idea for the urlPath to contain the bot token, as that makes it very difficult for outside
	// parties to find the update endpoint (which would allow them to inject their own updates).
	err = updater.StartWebhook(b, "telegram/"+token, webhookOpts)
	if err != nil {
		panic("failed to start webhook: " + err.Error())
	}

	err = updater.SetAllBotWebhooks(appUrl, &gotgbot.SetWebhookOpts{
		MaxConnections:     100,
		DropPendingUpdates: false,
		SecretToken:        webhookOpts.SecretToken,
	})
	if err != nil {
		panic("failed to set webhook: " + err.Error())
	}

	log.Printf("%s has been started...\n", b.User.Username)
	updater.Idle()
}
