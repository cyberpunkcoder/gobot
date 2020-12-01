package command

import (
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

func (p *purge) execute(commandMessage *discordgo.MessageCreate, session *discordgo.Session) {
	author := commandMessage.Author.ID
	mentionedMembers := commandMessage.Mentions
	commandChannel := commandMessage.ChannelID

	// Last 100 messages in command channel
	messages, err := session.ChannelMessages(commandChannel, 100, commandMessage.ID, "", "")

	// Return if unable to check messages
	if err != nil {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> unable to check messages.")
		return
	}

	session.ChannelMessageDelete(commandChannel, commandMessage.ID)
	purgeMessage, _ := session.ChannelMessageSend(commandChannel, "<@"+author+"> purging.")

	// If no user was mentioned in the command delete each of last 100 messages
	if len(mentionedMembers) == 0 {
		if strings.Contains(commandMessage.Content, " ") {
			session.ChannelMessageSend(commandChannel, "<@"+author+"> "+p.wrongFormat())
		} else {
			for _, message := range messages {
				// No user was mentioned, delete every message
				session.ChannelMessageDelete(commandChannel, message.ID)
			}
		}
	} else {
		for _, message := range messages {
			// For each of the users mentioned in the command
			for _, member := range mentionedMembers {
				if message.Author.ID == member.ID {
					session.ChannelMessageDelete(commandChannel, message.ID)
				}
			}
		}
	}

	// Delete purging status message
	session.ChannelMessageDelete(commandChannel, purgeMessage.ID)
}
