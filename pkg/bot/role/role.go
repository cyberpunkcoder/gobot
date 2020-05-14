package role

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// ReactionRoleCatagories catagories, and nameid emojiid relation for each reaction role
var (
	ReactionRoleCatagories reactionRoleCatagories
)

type reactionRoleCatagories []struct {
	Name string `json:"name"`
	Role []struct {
		RoleID  string `json:"roleid"`
		EmojiID string `json:"emojiid"`
	} `json:"role"`
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

	err = json.Unmarshal(file, &ReactionRoleCatagories)

	// Return if there was an error unmarshaling reactionrroles.json
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
