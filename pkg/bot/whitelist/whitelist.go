package whitelist

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	// BinPath is the location of the executable binary
	BinPath string

	// Whitelists that members will be muted for messaging
	Whitelists []whitelist

	whitelistsFile = "/json/whitelists.json"
)

// Filter phrase or word members will be muted for messaging
type whitelist struct {
	Text string `json:"text"`
}

// ContainsWiteList returns true if a message contains a whitelist
func ContainsWiteList(message *discordgo.MessageCreate, session *discordgo.Session) bool {
	for _, whitelist := range Whitelists {
		if strings.Contains(message.Content, whitelist.Text) && !message.Author.Bot {
			return true
		}
	}
	return false
}

// LoadWhitelists and whitelists alert from json file
func LoadWhitelists() error {
	log.Println("Loading Whitelists")

	file, err := ioutil.ReadFile(BinPath + whitelistsFile)

	// Create new json file if none exists, returns for other errors
	if err != nil {
		if strings.Contains(err.Error(), "no") && strings.Contains(err.Error(), "directory") {
			saveWhitelists()
			return nil
		}
		return err
	}

	err = json.Unmarshal(file, &Whitelists)

	// Return if there was an error unmarshaling json
	if err != nil {
		return err
	}
	return nil
}

// SaveWhitelist to json file
func SaveWhitelist(text string) error {
	for _, whitelist := range Whitelists {
		if whitelist.Text == text {
			// If the whitelist already exists return an error
			return errors.New("Filter already exists \"" + text + "\"")
		}
	}

	newFilter := whitelist{Text: text}
	Whitelists = append(Whitelists, newFilter)
	saveWhitelists()
	return nil
}

// RemoveWhitelist from json file
func RemoveWhitelist(text string) error {
	for i, whitelist := range Whitelists {
		if whitelist.Text == text {
			// Remove the whitelist from Whitelists
			Whitelists[i] = Whitelists[len(Whitelists)-1]
			Whitelists = Whitelists[:len(Whitelists)-1]

			saveWhitelists()
			return nil
		}
	}
	// If the whitelist was not found return an error
	return errors.New("Cannot find whitelist \"" + text + "\"")
}

func saveWhitelists() error {
	WhitelistsJSON, err := json.MarshalIndent(Whitelists, "", " ")

	// Return if there was an error marshaling Whitelists
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = ioutil.WriteFile(BinPath+whitelistsFile, WhitelistsJSON, 0644)

	// Return if there was an error writing to json file
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
