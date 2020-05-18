package role

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	// ReactionRoleCatagories nameid and emojiid relation for each reaction role
	ReactionRoleCatagories reactionRoleCatagories

	// ReactionRoleMessages messageid of all messages users react too to get roles
	ReactionRoleMessages reactionRoleMessages
)

type reactionRoleCatagories []struct {
	Name string `json:"name"`
	Role []struct {
		RoleID  string `json:"roleid"`
		EmojiID string `json:"emojiid"`
	} `json:"role"`
}

type reactionRoleMessages []string

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

// LoadReactionRoleMessages asd
func LoadReactionRoleMessages() error {

	log.Println("Loading reaction role messages")

	file, err := ioutil.ReadFile("../../config/reactionrolemessages.json")

	// Return if there was an error reading reactionroles.json
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &ReactionRoleMessages)

	// Return if there was an error unmarshaling reactionrrolemessages.json
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

// SaveReactionRoleMessage saves id of reaction role message for checking when users react to any message
func SaveReactionRoleMessage(messageid string) error {

	ReactionRoleMessages = append(ReactionRoleMessages, messageid)
	roleMessagesJSON, err := json.Marshal(ReactionRoleMessages)

	// Return if there was an error marshaling reactionrolemessages.json
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = ioutil.WriteFile("../../config/reactionrolemessages.json", roleMessagesJSON, 0644)

	// Return if there was an error writing to reactionrolemessages.json
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
