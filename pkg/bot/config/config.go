package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	// DiscordToken token to access bot application on Discord
	DiscordToken string

	// CommandPrefix character to represent the beginning of a bot command
	CommandPrefix string

	// JoinRole role members are given after joining guild
	JoinRole string

	// Private variables
	config *configStruct
)

type configStruct struct {
	DiscordToken  string `json:"DiscordToken"`
	CommandPrefix string `json:"CommandPrefix"`
	JoinRole      string `json:"JoinRole"`
}

// LoadConfig read and load the config.json file
func LoadConfig() error {
	log.Println("Loading config")

	file, err := ioutil.ReadFile("../../json/config.json")

	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	DiscordToken = config.DiscordToken
	CommandPrefix = config.CommandPrefix
	JoinRole = config.JoinRole

	return nil
}
