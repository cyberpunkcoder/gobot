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
	// Get keyword from string by removing everything after first space
	keyword := strings.Split(s, " ")[0]
	for _, whitelist := range Whitelists {
		if whitelist.Text == keyword {
			// Violates whitelist
			return false
		}
	}
	return true
}

// LoadWhitelists and whitelists alert from json file
func LoadWhitelists() error {
	log.Println("Loading Whitelists")

	file, err := ioutil.ReadFile(BinPath + whitelistsFile)
	if err != nil {
		// Create new whitelists.json file if none exists, returns for other errors
		if strings.Contains(err.Error(), "no") && strings.Contains(err.Error(), "directory") {
			return saveWhitelists()
		}
		return err
	}
	return json.Unmarshal(file, &Whitelists)
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
	return saveWhitelists()
}

// RemoveWhitelist from json file
func RemoveWhitelist(text string) error {
	for i, whitelist := range Whitelists {
		if whitelist.Text == text {
			// Remove the whitelist from Whitelists
			Whitelists[i] = Whitelists[len(Whitelists)-1]
			Whitelists = Whitelists[:len(Whitelists)-1]
			return saveWhitelists()
		}
	}
	// If the whitelist was not found return an error
	return errors.New("Cannot find whitelist \"" + text + "\"")
}

func saveWhitelists() error {
	WhitelistsJSON, err := json.MarshalIndent(Whitelists, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(BinPath+whitelistsFile, WhitelistsJSON, 0644)
	if err != nil {
		return err
	}
	return nil
}
