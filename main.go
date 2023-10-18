package main

import (
	"log"
	"os"

	"reminderbot/bot"
)

func main() {
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("Must set Discord token as env variable: BOT_TOKEN")
	}
	reminderBotChannel, ok := os.LookupEnv("CHANNEL_ID")
	if !ok {
		log.Fatal("Must set channel ID token as env variable: CHANNEL_ID")
	}
	reminderBotMessage, ok := os.LookupEnv("MESSAGE")
	if !ok {
		log.Fatal("Must set channel ID token as env variable: MESSAGE")
	}

	// Start the bot
	bot.ReminderBotChannel = reminderBotChannel
	bot.ReminderBotMessage = reminderBotMessage
	bot.BotToken = botToken
	bot.Run()
}
