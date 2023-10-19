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
	// Set up a discord session with our token
	discord, err0 := discordgo.New("Bot " + BotToken)
	if err0 != nil {
		log.Fatal(err0)
	}
	// Open and cleanly close the connection with discord
	err1 := discord.Open()
	if err1 != nil {
		log.Fatal(err1)
	}
	defer discord.Close()
	for {
		// Get the current day of the month
		CurrentDay, CurrentHour := GetTime()
		// Actually send or dont send a message if the date is the 1st (soon whatever the user sets it to)
		if (CurrentDay == ReminderDay) && (CurrentHour == ReminderHour) {
			log.Println("it is the configured date and hour so we are running")
			discord.ChannelMessageSend(ReminderBotChannel, ReminderBotMessage)
		} else {
			log.Println("it is currently", CurrentDay, "and", CurrentHour, "and not", ReminderDay, "and", ReminderHour, "so we are not running")
		}
		time.Sleep(time.Hour)
	}
}

func GetTime() (int, int) {
	t := time.Now()
	hour := t.Hour()
	day := t.Day()
	log.Println("Reading", hour, "as current hour and", day, "as current day")
	return hour, day
}
