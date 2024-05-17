package modules

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/botsgalaxy/render-go-telegram-bot-template/internal/database"
)

func Start(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	chatId := ctx.EffectiveChat.Id
	text := fmt.Sprintf(`👋 Heya <b>%s</b>,
	
	✅ This is a sample bot running on render.com`, user.FirstName)
	button := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				gotgbot.InlineKeyboardButton{Text: "</> Developer", Url: "https://t.me/primeakash"},
				gotgbot.InlineKeyboardButton{Text: "📢 Update Channel", Url: "https://t.me/BotsgalaxyOfficial"},
			},
			{
				gotgbot.InlineKeyboardButton{Text: "#️⃣ Source Code", Url: "https://github.com/botsgalaxy/render-go-telegram-bot-template"},
			},
		},
	}

	_, err := b.SendMessage(
		chatId,
		text,
		&gotgbot.SendMessageOpts{
			ParseMode:   "html",
			ReplyMarkup: button,
		},
	)
	telegramUser := database.TelegramUser{
		UserId:       user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		LanguageCode: user.LanguageCode,
		IsPremium:    user.IsPremium,
	}
	telegramUser.Create()

	return err

}
