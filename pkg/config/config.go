package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	// DiscordToken token to access bot application on Discord
	DiscordToken string

	// CommandPrefix character to represent the beginning of a bot command
	CommandPrefix string

	// Private variables
	config *configStruct
)

type configStruct struct {
	DiscordToken  string `json:"DiscordToken"`
	CommandPrefix string `json:"CommandPrefix"`
}

// ReadConfig read and load the config.json file from config directory in root of the project
func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := ioutil.ReadFile("../../config/config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	DiscordToken = config.DiscordToken
	CommandPrefix = config.CommandPrefix

	return nil
}
