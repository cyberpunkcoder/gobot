/*
Gobot, a robust expandable bot that interfaces with Discord go API!
Copyright (C) 2020 github.com/cyberpunkprogrammer

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
	filter.BinPath = BinPath

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
	}

	// Load roles assigned by reaction
	err = reactionrole.LoadMessages()

	// Return if error loading reactionrroles.json
	if err != nil {
		log.Println("Error loading reaction role messages,", err)
	}

	// Load filters to be muted for
	err = filter.LoadFilters()

	// Return if error loading filters.json
	if err != nil {
		log.Println("Error loading filters,", err)
	}

	bot.Start()
}
