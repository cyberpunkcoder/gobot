package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
)

type welcome struct {
	command
}

func init() {
	welcome := welcome{command{
		name:        "welcome",
		parameters:  "(optional: @member)",
		description: "Sends welcome message in channel or addressed to member.",
		permissions: []int{discordgo.PermissionManageMessages},
	}}
	executables = append(executables, &welcome)
}

func (h *welcome) execute(message *discordgo.MessageCreate, session *discordgo.Session) error {
	if config.WelcomeMessage == "" {
		session.ChannelMessageSend(message.ChannelID, "<@"+message.Author.ID+"> welcome message not set.")
		return nil
	}
	session.ChannelMessageDelete(message.ChannelID, message.ID)

	mentionedMembers := message.Mentions
	// If no members mentioned, send welcome message to channel
	if len(mentionedMembers) == 0 {
		session.ChannelMessageSend(message.ChannelID, config.WelcomeMessage)
	}
	// Send welcome message to mentioned members
	for _, member := range mentionedMembers {
		session.ChannelMessageSend(message.ChannelID, "<@"+member.ID+">, "+config.WelcomeMessage)
	}
	return nil
}
