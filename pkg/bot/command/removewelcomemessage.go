package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
)

type removeWelcomeMessage struct {
	command
}

func init() {
	removeWelcomeMessage := removeWelcomeMessage{command{
		name:        "removewelcomemessage",
		description: "Removes welcome message bot sends members when they join.",
		permissions: []int{discordgo.PermissionManageRoles},
	}}
	executables = append(executables, &removeWelcomeMessage)
}

func (a *removeWelcomeMessage) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	author := message.Author.ID
	channel := message.ChannelID

	if config.WelcomeMessage == "" {
		session.ChannelMessageSend(channel, "<@"+author+"> welcome message not set.")
		return
	}

	config.SaveWelcomeMessage("")
	session.ChannelMessageSend(channel, "<@"+author+"> welcome message removed.")
}
