package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
)

type removeJoinRole struct {
	command
}

func init() {
	removeJoinRole := removeJoinRole{command{
		name:        "removejoinrole",
		description: "Removes role bot gives members when they join.",
		permissions: []int{discordgo.PermissionManageRoles},
	}}
	executables = append(executables, &removeJoinRole)
}

func (a *removeJoinRole) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	author := message.Author.ID
	channel := message.ChannelID

	if config.JoinRole == "" {
		session.ChannelMessageSend(channel, "<@"+author+"> join role not set.")
		return
	}

	config.SaveJoinRole("")
	session.ChannelMessageSend(channel, "<@"+author+"> join role removed.")
}
