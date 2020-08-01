package reactionrole

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

var (
	// Catagories nameid and emojiid relation for each reaction role
	Catagories Catagory

	// Messages messageid of all messages users react too to get roles
	Messages message

	reactionRoleFile         = "../../json/role/reactionroles.json"
	reactionRoleMessagesFile = "../../json/role/reactionrolemessages.json"
)

// Catagory of reaction roles each containing a number of reaction roles
type Catagory []struct {
	Name string `json:"name"`
	Role []Role `json:"role"`
}

// Role that contains a role ID and associated emoji struct
type Role struct {
	ID    string `json:"id"`
	Emoji Emoji  `json:"emoji"`
}

// Emoji that contains a prefix for an animated emoji and an associated name and ID
type Emoji struct {
	Prefix string `json:"prefix"`
	Name   string `json:"name"`
	ID     string `json:"id"`
}

// Message users react too to obtain roles
type message []string

// LoadRoles loads roles and their associated emoji from reactionroles.json
func LoadRoles() error {

	log.Println("Loading reaction roles")

	file, err := ioutil.ReadFile(reactionRoleFile)

	// Return if there was an error reading reactionroles.json
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &Catagories)

	// Return if there was an error unmarshaling reactionrroles.json
	if err != nil {
		return err
	}

	return nil
}

// SaveRole saves reaction role to appear in the reaction role message list
func SaveRole(catagoryName string, role Role) {

	// Add role to catagory with name catagoryName
	for i, catagory := range Catagories {
		if catagory.Name == catagoryName {
			Catagories[i].Role = append(catagory.Role, role)
			saveRoles()
			return
		}
	}

	// If no catagory was found under the name catagoryName, create a new one
	newCatagory := Catagory{
		{
			Name: catagoryName,
			Role: []Role{role},
		},
	}

	Catagories = append(Catagories, newCatagory...)
	saveRoles()
}

// RemoveRole removes a role from the reaction role menu
func RemoveRole(roleID string) error {
	for i, catagory := range Catagories {
		for j, reactionRole := range catagory.Role {

			// If the role id is found
			if reactionRole.ID == roleID {

				// Remove the role from the catagory
				Catagories[i].Role[j] = catagory.Role[len(catagory.Role)-1]
				Catagories[i].Role = catagory.Role[:len(catagory.Role)-1]

				// If catagory is empty, remove it
				if len(Catagories[i].Role) == 0 {
					Catagories[i] = Catagories[len(Catagories)-1]
					Catagories = Catagories[:len(Catagories)-1]
				}

				// Save changes to reactionrroles.json
				saveRoles()
				return nil
			}
		}
	}

	// If the role was not found return an error.
	return errors.New("Cannot find reaction role \"" + roleID + "\"")
}

// LoadMessages loads messages ids that users react too to obtain roles
func LoadMessages() error {

	log.Println("Loading reaction role messages")

	file, err := ioutil.ReadFile(reactionRoleMessagesFile)

	// Return if there was an error reading reactionroles.json
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &Messages)

	// Return if there was an error unmarshaling reactionrrolemessages.json
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

// SaveMessage saves id of reaction role message for checking when users react to any message
func SaveMessage(messageid string) error {

	Messages = append(Messages, messageid)
	roleMessagesJSON, err := json.MarshalIndent(Messages, "", " ")

	// Return if there was an error marshaling reactionrolemessages.json
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = ioutil.WriteFile(reactionRoleMessagesFile, roleMessagesJSON, 0644)

	// Return if there was an error writing to reactionrolemessages.json
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

// saveRoles saves the current state of the roles to reactionroles.json
func saveRoles() error {
	roleMessagesJSON, err := json.MarshalIndent(Catagories, "", " ")

	// Return if there was an error marshaling reactionrolemessages.json
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = ioutil.WriteFile(reactionRoleFile, roleMessagesJSON, 0644)

	// Return if there was an error writing to reactionrolemessages.json
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
