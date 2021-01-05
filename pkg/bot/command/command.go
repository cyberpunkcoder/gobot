package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
)

var (
	executables []executable
)

type executable interface {
	getName() string
	getUsage() string
	getDescription() string
	getPermisssions() []int
	execute(message *discordgo.MessageCreate, session *discordgo.Session)
}

type command struct {
	executable
	name        string
	parameters  string
	description string
	permissions []int
}

// Execute a command
func Execute(message *discordgo.MessageCreate, session *discordgo.Session) {

	// Get keyword fom message text by removing CommandPrefix and everything after first space
	keyword := strings.Replace(message.Content, config.CommandPrefix, "", -1)
	keyword = strings.Split(keyword, " ")[0]

	// Look through all commands to see if there is a match
	found := false
	for _, e := range executables {
		if e.getName() == keyword {
			found = true
			author := message.Author.ID
			channel := message.ChannelID

			botPermissions, err := session.State.UserChannelPermissions(session.State.User.ID, channel)

			if err != nil {
				session.ChannelMessageSend(channel, "<@"+author+"> unable to check my permissions.")
				return
			}

			var missingBotPermissions []int
			for _, p := range e.getPermisssions() {
				if botPermissions&p == 0 {
					missingBotPermissions = append(missingBotPermissions, p)
				}
			}

			if len(missingBotPermissions) > 0 {
				reply := "<@" + message.Author.ID + "> I don't have permission to execute this command."

				/*
					for _, p := range missingPermissions {
						TODO: List permissions missing
					}
				*/

				session.ChannelMessageSend(message.ChannelID, reply)
				return
			}

			userPermissions, err := session.State.UserChannelPermissions(author, channel)

			if err != nil {
				session.ChannelMessageSend(channel, "<@"+author+"> unable to check your permissions.")
				return
			}

			var missingUserPermissions []int
			for _, p := range e.getPermisssions() {
				if userPermissions&p == 0 {
					missingUserPermissions = append(missingUserPermissions, p)
				}
			}

			if len(missingUserPermissions) > 0 {
				reply := "<@" + message.Author.ID + "> you don't have permission to execute this command."

				/*
					for _, p := range missingPermissions {
						TODO: List permissions missing
					}
				*/

				session.ChannelMessageSend(message.ChannelID, reply)
				return
			}

			e.execute(message, session)
			return
		}
	}

	// If no matching command was found
	if !found {
		session.ChannelMessageSend(message.ChannelID, "<@"+message.Author.ID+"> unknown command, try **!help**.")
	}
}

func (c *command) getName() string {
	return c.name
}

func (c *command) getUsage() string {
	return config.CommandPrefix + c.name + " " + c.parameters
}

func (c *command) getDescription() string {
	return c.description
}

func (c *command) getPermisssions() []int {
	return c.permissions
}

func (c *command) wrongFormat() string {
	return "incorrectly formatted command, try **" + c.getUsage() + "**."
}
