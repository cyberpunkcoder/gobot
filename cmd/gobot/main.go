package main

import (
	"fmt"
	"log"

	"github.com/cyberpunkprogrammer/gobot/pkg/bot"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/reactionrole"
)

func main() {

	// Load configuration for bot and api
	err := config.LoadConfig()

	// Return if error loading config.json
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Load roles assigned by reaction
	err = reactionrole.LoadRoles()

	// Return if error loading reactionrroles.json
	if err != nil {
		log.Println("Error loading reaction roles,", err)
		return
	}

	// Load roles assigned by reaction
	err = reactionrole.LoadMessages()

	// Return if error loading reactionrroles.json
	if err != nil {
		log.Println("Error loading reaction role messages,", err)
		return
	}

	bot.Start()
}
