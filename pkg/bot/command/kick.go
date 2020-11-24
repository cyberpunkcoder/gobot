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
		description: "Kicks a mentioned member",
		permissions: []int{discordgo.PermissionKickMembers},
	}}

	executables = append(executables, &kick)
}

func (k *kick) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	author := message.Author.ID
	mentionedMembers := message.Mentions
	commandChannel := message.ChannelID

	// Return if no users mentioned in command
	if len(mentionedMembers) <= 0 {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> you must mention a user to kick.")
		return
	}

	// Kick all mentioned users
	for _, member := range mentionedMembers {
		if member.ID == author {
			session.ChannelMessageSend(commandChannel, "<@"+member.ID+"> you cannot kick yourself.")
		} else {
			session.ChannelMessageSend(commandChannel, "<@"+member.ID+"> was kicked.")
			session.GuildMemberDelete(message.GuildID, member.ID)
		}
	}
}
