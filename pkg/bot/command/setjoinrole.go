package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
)

type setJoinRole struct {
	command
}

func init() {
	setJoinRole := setJoinRole{command{
		name:        "setjoinrole",
		parameters:  "@role",
		description: "Sets a role bot gives members when they join.",
		permissions: []int{discordgo.PermissionManageRoles},
	}}
	executables = append(executables, &setJoinRole)
}

func (a *setJoinRole) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	author := message.Author.ID
	channel := message.ChannelID
	mentionedRoles := message.MentionRoles

	if len(mentionedRoles) == 1 {
		if config.JoinRole == mentionedRoles[0] {
			session.ChannelMessageSend(channel, "<@"+author+"> join role already set to that.")
			return
		}
		session.ChannelMessageSend(channel, "<@"+author+"> join role set.")
		config.SaveJoinRole(mentionedRoles[0])
		return
	}
	session.ChannelMessageSend(channel, "<@"+author+"> "+a.wrongFormat())
}
