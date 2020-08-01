package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
)

// Start the Discord bot
func Start() {

	// Create new bot session
	botSession, err := discordgo.New("Bot " + config.DiscordToken)

	// Check if bot was created
	if err != nil {
		log.Println("Error creating Discord session,", err)
		return
	}

	// Sever events
	botSession.AddHandler(ready)
	botSession.AddHandler(guildMemberAdd)
	botSession.AddHandler(messageCreate)
	botSession.AddHandler(messageReactionAdd)
	botSession.AddHandler(messageReactionRemove)

	err = botSession.Open()

	// Check if Discord connection was made
	if err != nil {
		log.Println("Error opening Discord connection,", err)
		return
	}

	// Run until CTRL-C or other term signal is received
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session
	botSession.Close()
}