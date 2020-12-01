package command

import (
	"github.com/bwmarrin/discordgo"
)

type kick struct {
	command
}

func init() {

	kick := kick{command{
		name:        "kick",
		parameters:  "@member",
		description: "Kicks a mentioned member.",
		permissions: []int{discordgo.PermissionKickMembers},
	}}

	executables = append(executables, &kick)
}

func (k *kick) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	author := message.Author.ID
	mentionedMembers := message.Mentions
	channel := message.ChannelID

	// Return if no users mentioned in command
	if len(mentionedMembers) <= 0 {
		session.ChannelMessageSend(channel, "<@"+author+"> "+k.wrongFormat())
		return
	}

	// Kick all mentioned users
	for _, member := range mentionedMembers {
		if member.ID == author {
			session.ChannelMessageSend(channel, "<@"+member.ID+"> you cannot kick yourself.")
		} else {
			session.ChannelMessageSend(channel, "<@"+member.ID+"> was kicked.")
			session.GuildMemberDelete(message.GuildID, member.ID)
		}
	}
}
