package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// подключаемся к боту с помощью токена
	// bot, err := tgbotapi.NewBotAPI("1281304074:AAFXLPttlX0sm92QxHdTFlyacer9Hapf_Ic")
	bot, err := tgbotapi.NewBotAPI("368741635:AAF_a7JaWZri88Hn4CJCqPkKZLq9UF49qA4")

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		var m *tgbotapi.Message
		var msg tgbotapi.MessageConfig

		if update.CallbackQuery != nil {
			m = update.CallbackQuery.Message
			msg = tgbotapi.NewMessage(m.Chat.ID, "")

			switch update.CallbackQuery.Data {
			case "begin":
				msg.Text = "Выбери из списка предметы, по которым ты хотел бы решать задачи"
				msg.ReplyMarkup = newPredmetKeyboard()
				bot.Send(msg)
			case "math":
				msg.Text = "Математика"
				bot.Send(msg)
			case "phis":
				msg.Text = "Физика"
				bot.Send(msg)
			case "it":
				msg.Text = "Информатика"
				bot.Send(msg)
			}
			continue
		} else {
			m = update.Message
			msg = tgbotapi.NewMessage(m.Chat.ID, "")
			msg.ParseMode = "markdown"
		}

		if m.IsCommand() {
			switch m.Command() {
			case "start":
				msg.Text = fmt.Sprintf("*Привет, %s!*\n\nЗдесь ты можешь каждый день решать интересные задачи и зарабатывать дополнительные баллы для поступления в КЛШ-2021.", m.From.FirstName)
				msg.ReplyMarkup = newWelcomeKeyboard()
				bot.Send(msg)
			case "key":
				msg.ReplyMarkup = newKeyboard()
				msg.Text = "Клавиатура"
				bot.Send(msg)
			case "photo":
				photo := tgbotapi.NewPhotoUpload(m.Chat.ID, "a_d21fbd44.jpg")
				m, e := bot.Send(photo)
				fmt.Println(m)
				fmt.Println(e)
			}
		} else if m.Text != "" {
			fmt.Println()
			fmt.Println(m.Chat)
			fmt.Println()
			msg.Text = "✔️ " + m.Text
			bot.Send(msg)
		} else {
			continue
		}
	}
}

func newKeyboard() (keyb tgbotapi.ReplyKeyboardMarkup) {
	btn1 := tgbotapi.NewKeyboardButton("Button 1")
	btn2 := tgbotapi.NewKeyboardButtonLocation("Location")
	btn3 := tgbotapi.NewKeyboardButton("Button 3")
	btn4 := tgbotapi.NewKeyboardButton("Button 4")
	row1 := tgbotapi.NewKeyboardButtonRow(btn1, btn2)
	row2 := tgbotapi.NewKeyboardButtonRow(btn3, btn4)
	keyb = tgbotapi.NewReplyKeyboard(row1, row2)
	return
}

func newWelcomeKeyboard() (keyb tgbotapi.InlineKeyboardMarkup) {
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Выбрать предметы", "begin")
	row1 := tgbotapi.NewInlineKeyboardRow(btn1)
	keyb = tgbotapi.NewInlineKeyboardMarkup(row1)
	return
}

func newPredmetKeyboard() (keyb tgbotapi.InlineKeyboardMarkup) {
	btn1 := tgbotapi.NewInlineKeyboardButtonData("Математика", "math")
	btn2 := tgbotapi.NewInlineKeyboardButtonData("Физика", "phis")
	btn3 := tgbotapi.NewInlineKeyboardButtonData("Информатика", "it")
	row1 := tgbotapi.NewInlineKeyboardRow(btn1, btn2, btn3)
	keyb = tgbotapi.NewInlineKeyboardMarkup(row1)
	return
}
