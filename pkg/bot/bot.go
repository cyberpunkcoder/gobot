package bot

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
)

const healthCheckPort = 8080

// Start the Discord bot
func Start() error {
	log.Println("Starting bot")
	dg, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		return err
	}

	log.Println("Serving events")
	// Serve events
	dg.AddHandler(ready)
	dg.AddHandler(guildMemberAdd)
	dg.AddHandler(messageCreate)
	dg.AddHandler(messageReactionAdd)
	dg.AddHandler(messageReactionRemove)

	// Gather all intents
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	log.Println("Opening bot session")
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

func ServeHealthCheck() error {
	log.Println("Serving health check")
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Health check served")
		w.WriteHeader(http.StatusOK)
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", healthCheckPort), nil)
}
