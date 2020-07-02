package main

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI("1281304074:AAFXLPttlX0sm92QxHdTFlyacer9Hapf_Ic")

	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		cmd := ""

		if update.CallbackQuery != nil {
			m := update.CallbackQuery.Message
			// log.Println(update.CallbackQuery.Data)
			switch update.CallbackQuery.Data {
			case "time":
				cmd = "time"
			}
		} else {
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}
			m := update.Message
			cmd = m.Command()
		}

		// log.Println(update.Message)

		msg := tgbotapi.NewMessage(m.Chat.ID, "")

		if cmd != "" {
			switch m.Command() {
			case "start":
				msg.Text = fmt.Sprintf("Привет, %s!", m.From.FirstName)
				keyboard := tgbotapi.InlineKeyboardMarkup{}
				var row []tgbotapi.InlineKeyboardButton
				btn := tgbotapi.NewInlineKeyboardButtonData("Время в Красноярске", "time")
				row = append(row, btn)
				keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
				msg.ReplyMarkup = keyboard
			case "time":
				t := time.Now()
				msg.Text = fmt.Sprintf("Сейчас в Красноярске %d:%d:%d", t.Hour(), t.Minute(), t.Second())
			}
		} else {
			msg.Text = m.Text
		}

		bot.Send(msg)
	}
}
