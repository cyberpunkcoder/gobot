package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cyberpunkprogrammer/gobot/pkg/bot"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/reactionrole"
)

// BinPath is the location of the executable binary
var BinPath string

func main() {
	BinPath, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		log.Println(err)
		return
	}

	// Create directory nstructure if it does not already exist
	_ = os.MkdirAll(BinPath+"/json/role", 0777)

	// Set relative path to packages that access json files
	config.BinPath = BinPath
	reactionrole.BinPath = BinPath

	// Load configuration for bot and api
	err = config.LoadConfig()

	// Return if error loading config.json
	if err != nil {
		if strings.Contains(err.Error(), "no") && strings.Contains(err.Error(), "directory") {
			fmt.Println("Configuration not found")
			config.Create()
		} else {
			log.Println(err)
			return
		}
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
