package bot

import (
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Pull these from the other function...
var (
	ReminderBotChannel string
	ReminderBotMessage string
	BotToken           string
)

func Run() {
	// Get the current day of the month
	day := GetTime()
	// Set up a discord session with our token
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
	}
	// Open and cleanly close the connection with discord
	discord.Open()
	defer discord.Close()
	// Actually send or dont send a message if the date is the 1st (soon whatever the user sets it to)
	if day == 1 {
		log.Println("it is the first so we are running")
		discord.ChannelMessageSend(ReminderBotChannel, ReminderBotMessage)
	}
	if day > 1 {
		log.Println("it is not the 1st so there's no need to run right now")
	}
	// This is here so cron doesn't assume the process didn't complete properly
	os.Exit(0)
}

func GetTime() int {
	// Gets the time...
	d := time.Now()
	// Changes it to just the day of the month...
	day := d.Day()
	// Log it for debugging
	log.Println(day)
	// and return
	return day
}
