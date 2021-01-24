package whitelist

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strings"
)

var (
	// BinPath is the location of the executable binary
	BinPath string

	// Whitelists that the bot will ignore
	Whitelists []whitelist

	whitelistsFile = "/json/whitelists.json"
)

// whitelist the bot will ignore
type whitelist struct {
	Text string `json:"text"`
}

// Check returns false if a message contains a whitelist, true if it does not
func Check(s string) bool {
	for _, whitelist := range Whitelists {
		if strings.Contains(s, whitelist.Text) {
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
