package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/config"
)

func Run() {

	discord, err := discordgo.New("Bot " + config.DiscordToken)

	// Check if Discord session was created
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	discord.AddHandler(ready)
	discord.AddHandler(messageCreate)

	err = discord.Open()

	// Check if Discord connection was made
	if err != nil {
		fmt.Println("Error opening Discord connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session
	discord.Close()
}

// This function will be called (due to AddHandler above) every time the bot sucessfully logs in
func ready(s *discordgo.Session, r *discordgo.Ready) {

	// Display name of bot to user after login successful
	fmt.Println("Logged in as", s.State.User, "Press CTRL-C to exit")
}

// This function will be called (due to AddHandler above) every time a new message is created on any channel that the autenticated bot has access to
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
