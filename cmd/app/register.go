package main

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/botsgalaxy/render-go-telegram-bot-template/internal/modules"
)

func addHandlers(d *ext.Dispatcher) {
	d.AddHandler(handlers.NewCommand("start", modules.Start))
	d.AddHandler(handlers.NewCommand("ping", modules.Ping))
}
