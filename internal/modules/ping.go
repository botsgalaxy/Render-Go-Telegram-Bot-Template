package modules

import (
	"fmt"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Ping(b *gotgbot.Bot, ctx *ext.Context) error {
	chatId := ctx.EffectiveChat.Id
	msgId := ctx.EffectiveMessage.MessageId

	start := time.Now()
	msg, err := b.SendMessage(
		chatId,
		"! Pong...",
		&gotgbot.SendMessageOpts{
			ReplyParameters: &gotgbot.ReplyParameters{
				MessageId: msgId,
				ChatId:    chatId,
			},
		},
	)
	if err != nil {
		return err
	}
	end := time.Now()
	duration := end.Sub(start).Microseconds() / 1000
	_, _, err = msg.EditText(
		b,
		fmt.Sprintf("! Pong... %d ms", duration),
		nil,
	)
	return err
}
