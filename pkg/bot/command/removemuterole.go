package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
)

type removeMuteRole struct {
	command
}

func init() {
	removeMuteRole := removeMuteRole{command{
		name:        "removemuterole",
		description: "Removes role bot gives members to mute them.",
		permissions: []int{discordgo.PermissionManageRoles},
	}}
	executables = append(executables, &removeMuteRole)
}

func (a *removeMuteRole) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	author := message.Author.ID
	channel := message.ChannelID

	if config.MuteRole == "" {
		session.ChannelMessageSend(channel, "<@"+author+"> mute role not set.")
		return
	}

	config.SaveMuteRole("")
	session.ChannelMessageSend(channel, "<@"+author+"> mute role removed.")
}
