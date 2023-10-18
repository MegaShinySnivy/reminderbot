package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"reminderbot/bot"
)

func main() {
	// This gets all of our bot's required information.
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
		log.Fatal("Must set message as env variable: MESSAGE")
	}
	ReminderDayString, ok := os.LookupEnv("REMINDER_DAY")
	if !ok {
		log.Fatal("Must set day as env variable: REMINDER_DAY. Example: 13, 1, 8")
	}
	ReminderHourString, ok := os.LookupEnv("REMINDER_HOUR")
	if !ok {
		log.Fatal("Must set hour as env variable: REMINDER_HOUR. example: 13, 1, 8")
	}

	// Start the bot
	// This changes the string we get from REMINDER_DAY into an int so we can do a comparison later
	bot.ReminderDay = convertDay(ReminderDayString)
	bot.ReminderHour = convertHour(ReminderHourString)
	bot.ReminderBotChannel = reminderBotChannel
	bot.ReminderBotMessage = reminderBotMessage
	bot.BotToken = botToken
	for {
		bot.Run()
		// wait to check again until one hour has passed
		time.Sleep(60 * time.Minute)
	}
}

func convertDay(ReminderDayString string) int {
	ReminderDay, err := strconv.Atoi(ReminderDayString)
	if err != nil {
		panic(err)
	}
	return ReminderDay
}

func convertHour(ReminderHourString string) int {
	ReminderHour, err := strconv.Atoi(ReminderHourString)
	if err != nil {
		panic(err)
	}
	return ReminderHour
}
