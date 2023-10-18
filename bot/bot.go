package bot

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Pull these from the other function...
var (
	ReminderBotChannel string
	ReminderBotMessage string
	BotToken           string
	ReminderDay        int
	ReminderHour       int
)

func Run() {
	// Get the current day of the month
	day := GetDay()
	hour := GetTime()
	// Set up a discord session with our token
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
	}
	// Open and cleanly close the connection with discord
	discord.Open()
	defer discord.Close()
	// Actually send or dont send a message if the date is the 1st (soon whatever the user sets it to)
	if (day == ReminderDay) && (hour == ReminderHour) {
		log.Println("it is the configured date and hour so we are running")
		discord.ChannelMessageSend(ReminderBotChannel, ReminderBotMessage)
	}
	if day != ReminderDay {
		log.Println("it is not the configured date so there's no need to run right now")
	}
	if hour != ReminderHour {
		log.Println("it is the configured date, but not time, so no need to run right now")
	}
}

func GetDay() int {
	// Gets the time...
	d := time.Now()
	// Changes it to just the day of the month...
	day := d.Day()
	// Log it for debugging
	log.Println(day)
	// and return
	return day
}

func GetTime() int {
	// Gets the time...
	t := time.Now()
	// Changes it to just the current hour...
	hour := t.Hour()
	// Log it for debugging
	log.Println(hour)
	// and return
	return hour
}
