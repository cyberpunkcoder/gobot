package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
<<<<<<< HEAD
=======
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/role"
>>>>>>> 966c18fce780f4ff09ed19326f064994d2bce5b2
	"github.com/cyberpunkprogrammer/gobot/pkg/config"
)

// Start the Discord bot
func Start() {

	// Load roles assigned by reaction
<<<<<<< HEAD
	err := LoadReactionRoles()
=======
	err := role.LoadReactionRoles()
>>>>>>> 966c18fce780f4ff09ed19326f064994d2bce5b2

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
