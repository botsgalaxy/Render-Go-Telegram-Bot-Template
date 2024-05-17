## Render-Go-Telegram-Bot-Template

### Features 

- Use Webhook for updates.
- Host on render.com.

### Environment Variables 

- `BOT_TOKEN`: Your telegram bot token.
- `PORT`: Your port where webhook will listen.
- `APP_URL` : Your domain for webhook.
- `WEBHOOK_SECRET`: Something Secure. 
- `DSN`: Your database DSN for GORM.

### Render 
- BuildCommand : `go build -tags netgo -ldflags '-s -w' -o app`
- StartCommand : `./app`



### Author 
- Github: @botsgalaxy
- Telegram: @primeakash