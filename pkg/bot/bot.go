package bot

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
)

// Start the Discord bot
func Start() error {

	// Create new bot session
	dg, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		return err
	}

	// Sever events
	dg.AddHandler(ready)
	dg.AddHandler(guildMemberAdd)
	dg.AddHandler(messageCreate)
	dg.AddHandler(messageReactionAdd)
	dg.AddHandler(messageReactionRemove)

	// Gather all intents
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	err = dg.Open()
	if err != nil {
		return err
	}

	// Run until CTRL-C or other term signal is received
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session
	dg.Close()

	return nil
}
