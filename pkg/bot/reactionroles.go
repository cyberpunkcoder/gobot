package bot

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// ReactionRoles are roles that users obtain by reacting to a message
var (
	ReactionRoles []reactionRole
)

type reactionRole struct {
	ID    string `json:"id"`
	Emoji string `json:"emoji"`
}

// LoadReactionRoles loads roles and their associated emoji from reactionroles.json
func LoadReactionRoles() error {

	log.Println("Loading reaction roles")

	file, err := ioutil.ReadFile("../../config/reactionroles.json")

	// Return if there was an error reading reactionroles.json
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &ReactionRoles)

	// Return if there was an error unmarshaling reactionrroles.json
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
