package bot

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

// Pull these from the other function...
var (
	ReminderBotChannel string
	ReminderBotMessage string
	BotToken           string
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
	// Send a message on run
	defer discord.Close()
	log.Println("it is the configured date and hour so we are sending the following message:", ReminderBotMessage)
	_, err2 := discord.ChannelMessageSend(ReminderBotChannel, ReminderBotMessage)
	if err2 != nil {
		log.Fatal(err2)
	}
	os.Exit(0)
}
