package bot

import (
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	ReminderBotChannel string
	ReminderBotMessage string
	BotToken           string
)

func Run() {
	day := GetTime()
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
	}

	discord.Open()
	defer discord.Close()

	if day == 1 {
		log.Println("it is the first so we are running")
		discord.ChannelMessageSend(ReminderBotChannel, ReminderBotMessage)
	}
	if day > 1 {
		log.Println("it is not the 1st so there's no need to run right now")
	}

	os.Exit(0)
}

func GetTime() int {
	d := time.Now()
	day := d.Day()
	log.Println(day)
	return day
}
