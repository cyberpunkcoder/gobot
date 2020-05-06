package bot

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/command"
	"github.com/cyberpunkprogrammer/gobot/pkg/config"
)

// Start the Discord bot
func Start() {

	// Create new bot session
	botSession, err := discordgo.New("Bot " + config.DiscordToken)

	// Check if bot was created
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	botSession.AddHandler(ready)
	botSession.AddHandler(messageCreate)

	err = botSession.Open()

	// Check if Discord connection was made
	if err != nil {
		fmt.Println("Error opening Discord connection,", err)
		return
	}

	// Run until CTRL-C or other term signal is received
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session
	botSession.Close()
}

// ready is called whenever the bot has successfully logged in
func ready(session *discordgo.Session, ready *discordgo.Ready) {

	// Display name of bot to user
	fmt.Println("Logged in as", session.State.User, "Press CTRL-C to exit")
}

// messageCreate is called whenever a message has been created
func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	if strings.HasPrefix(message.Content, config.CommandPrefix) {

		commandText := message.Content[1:]
		fmt.Println(commandText)

		// Ignore all messages created by the bot itself
		if message.Author.ID == session.State.User.ID {
			return
		}

		if commandText == "hello" {
			command.Hello(session, message)
		}
	}
}
