package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
)

type setMuteRole struct {
	command
}

func init() {
	setMuteRole := setMuteRole{command{
		name:        "setmuterole",
		parameters:  "@role",
		description: "Sets a role bot gives members to mute them.",
		permissions: []int{discordgo.PermissionManageRoles},
	}}
	executables = append(executables, &setMuteRole)
}

func (a *setMuteRole) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	author := message.Author.ID
	channel := message.ChannelID
	mentionedRoles := message.MentionRoles

	if len(mentionedRoles) == 1 {
		if config.MuteRole == mentionedRoles[0] {
			session.ChannelMessageSend(channel, "<@"+author+"> mute role already set to that.")
			return
		}
		session.ChannelMessageSend(channel, "<@"+author+"> mute role set.")
		config.SaveMuteRole(mentionedRoles[0])
		return
	}
	session.ChannelMessageSend(channel, "<@"+author+"> "+a.wrongFormat())
}
