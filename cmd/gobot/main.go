/*
Gobot, a robust expandable bot that interfaces with Discord go API!
Copyright (C) 2024 github.com/cyberpunkcoder

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cyberpunkprogrammer/gobot/pkg/bot"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/filter"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/reactionrole"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/whitelist"
)

// BinPath is the location of the executable binary
var BinPath string

func main() {
	BinPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("Error getting binary path: %v", err)
	}
	// Create directory structure if it does not already exist
	err = os.MkdirAll(BinPath+"/json/role", 0777)
	if err != nil {
		log.Fatalf("Error creating directory structure: %v", err)
	}

	// Set relative path to packages that access json files
	config.BinPath = BinPath
	reactionrole.BinPath = BinPath
	filter.BinPath = BinPath
	whitelist.BinPath = BinPath

	// Load configuration for bot and api
	err = config.LoadConfig()
	if err != nil {
		if strings.Contains(err.Error(), "no") && strings.Contains(err.Error(), "directory") {
			fmt.Println("Configuration not found")
			config.Create()
		} else {
			log.Fatalf("Error loading configuration: %v", err)
		}
	}

	// Load roles assigned by reaction
	err = reactionrole.LoadRoles()
	if err != nil {
		log.Printf("Error loading reaction roles: %v", err)
	}
	// Load roles assigned by reaction
	err = reactionrole.LoadMessages()
	if err != nil {
		log.Printf("Error loading reaction messages: %v", err)
	}
	// Load filters to be muted for
	err = filter.LoadFilters()
	if err != nil {
		log.Printf("Error loading filters: %v", err)
	}
	// Load whitelists bot will ignore
	err = whitelist.LoadWhitelists()
	if err != nil {
		log.Printf("Error loading whitelists: %v", err)
	}

	err = bot.ServeHealthCheck()
	if err != nil {
		log.Fatalf("Error serving health check: %v", err)
	}

	err = bot.Start()
	if err != nil {
		log.Fatalf("Error starting bot: %v", err)
	}
}
