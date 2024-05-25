package command

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type purge struct {
	command
}

func init() {
	purge := purge{command{
		name:        "purge",
		parameters:  "(optional: @member)",
		description: "Deletes last 100 messages from all or mentioned member.",
		permissions: []int{discordgo.PermissionManageMessages},
	}}
	executables = append(executables, &purge)
}

func (p *purge) execute(commandMessage *discordgo.MessageCreate, session *discordgo.Session) error {
	author := commandMessage.Author.ID
	mentionedMembers := commandMessage.Mentions
	channel := commandMessage.ChannelID

	// Last 100 messages in command channel
	messages, err := session.ChannelMessages(channel, 100, commandMessage.ID, "", "")
	if err != nil {
		log.Printf("Error getting messages: %v", err)
		session.ChannelMessageSend(channel, "<@"+author+"> unable to check messages.")
		return err
	}

	session.ChannelMessageDelete(channel, commandMessage.ID)
	purgeMessage, err := session.ChannelMessageSend(channel, "<@"+author+"> purging.")
	if err != nil {
		log.Printf("Error sending message: %v", err)
		return err
	}

	// If no user was mentioned in the command delete each of last 100 messages
	if len(mentionedMembers) == 0 {
		if strings.Contains(commandMessage.Content, " ") {
			session.ChannelMessageSend(channel, "<@"+author+"> "+p.wrongFormat())
		} else {
			for _, message := range messages {
				// No user was mentioned, delete every message
				session.ChannelMessageDelete(channel, message.ID)
			}
		}
	} else {
		for _, message := range messages {
			// For each of the users mentioned in the command
			for _, member := range mentionedMembers {
				if message.Author.ID == member.ID {
					session.ChannelMessageDelete(channel, message.ID)
				}
			}
		}
	}

	// Delete purging status message
	session.ChannelMessageDelete(channel, purgeMessage.ID)
	return nil
}
