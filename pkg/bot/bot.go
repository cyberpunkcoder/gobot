package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/role"
	"github.com/cyberpunkprogrammer/gobot/pkg/config"
)

// Start the Discord bot
func Start() {

	// Load roles assigned by reaction
	err := role.LoadReactionRoles()

	// Check if roles loaded
	if err != nil {
		log.Println("Error loading reaction roles,", err)
		return
	}

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
