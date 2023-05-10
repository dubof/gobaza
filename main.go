package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5880930536:AAH0U6DJbVC7q_2LFyubp9KCPKTHClu_6kE")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Запуск бота %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Text {
		case "Мяч⚽️":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите действие:")
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Шар"),
					tgbotapi.NewKeyboardButton("Ком"),
				),
			)
			bot.Send(msg)
		case "Капля 💧":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите действие:")
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Крупица"),
					tgbotapi.NewKeyboardButton("Частичка"),
				),
			)
			bot.Send(msg)
		case "Батон🥖":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите действие:")
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Хлеб"),
					tgbotapi.NewKeyboardButton("Булка"),
				),
			)
			bot.Send(msg)
		case "Смартфон📱":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите действие:")
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Телефон"),
					tgbotapi.NewKeyboardButton("Гаджет"),
				),
			)
			bot.Send(msg)
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Это полезный бот слов-синонимов!")
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Мяч⚽️"),
					tgbotapi.NewKeyboardButton("Капля 💧"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Батон🥖"),
					tgbotapi.NewKeyboardButton("Смартфон📱"),
				),
			)
			bot.Send(msg)
		}
	}
}
