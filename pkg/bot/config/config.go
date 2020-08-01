package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	// BinPath is the location of the executable binary
	BinPath string

	// DiscordToken token to access bot application on Discord
	DiscordToken string

	// CommandPrefix character to represent the beginning of a bot command
	CommandPrefix string

	// JoinRole role members are given after joining guild
	JoinRole string

	// configFile is the file which contians the program startup configuration
	configFile = "/json/config.json"
)

type configStruct struct {
	DiscordToken  string
	CommandPrefix string
	JoinRole      string
}

// LoadConfig read and load the config.json file
func LoadConfig() error {
	log.Println("Loading configuration")

	file, err := ioutil.ReadFile(BinPath + configFile)

	// Return if there was a problem reading the file
	if err != nil {
		return err
	}

	load := configStruct{}

	err = json.Unmarshal(file, &load)

	if err != nil {
		return err
	}

	DiscordToken = load.DiscordToken
	CommandPrefix = load.CommandPrefix
	JoinRole = load.JoinRole

	return nil
}

// SaveConfig saves the configuration to drive from memory
func SaveConfig() error {

	save := configStruct{
		DiscordToken:  DiscordToken,
		CommandPrefix: CommandPrefix,
		JoinRole:      JoinRole,
	}

	configJSON, err := json.MarshalIndent(save, "", " ")

	// Return if there was a problem marshaling config
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(BinPath+configFile, configJSON, 0644)

	// Return if there was an error writing to the file
	if err != nil {
		return err
	}

	return nil
}

// Create creates a new configuration by prompting the user
func Create() {
	fmt.Print("Create new configuration? (y/n): ")
	for {
		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')

		if err != nil {
			log.Println(err.Error())
			return
		}

		if response == "y\n" {

			// Obtain bot token from user and store into variable
			fmt.Print("Enter discord bot token: ")
			discordToken, err := reader.ReadString('\n')
			DiscordToken = strings.TrimRight(discordToken, "\n")

			if err != nil {
				log.Println(err)
				return
			}

			// Obtain command prefix from user and store into variable
			fmt.Print("Enter command prefix (For example '!'): ")
			commandPrefix, err := reader.ReadString('\n')
			CommandPrefix = strings.TrimRight(commandPrefix, "\n")

			if err != nil {
				log.Println(err)
				return
			}

			// Obtain role id from user and store into variable
			fmt.Print("Enter role ID users get after joining (blank for none): ")
			joinRole, err := reader.ReadString('\n')
			JoinRole = strings.TrimRight(joinRole, "\n")

			if err != nil {
				log.Println(err)
				return
			}

			SaveConfig()
			return
		}

		if response == "n\n" {
			return
		}
	}
}
