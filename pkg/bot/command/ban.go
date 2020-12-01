package command

import (
	"github.com/bwmarrin/discordgo"
)

type ban struct {
	command
}

func init() {

	ban := ban{command{
		name:        "ban",
		parameters:  "@member",
		description: "Bans a mentioned member.",
		permissions: []int{discordgo.PermissionBanMembers},
	}}

	executables = append(executables, &ban)
}

func (b *ban) execute(message *discordgo.MessageCreate, session *discordgo.Session) {

	author := message.Author.ID
	mentionedMembers := message.Mentions
	channel := message.ChannelID

	// Return if no users mentioned in command
	if len(mentionedMembers) <= 0 {
		session.ChannelMessageSend(channel, "<@"+author+"> "+b.wrongFormat())
		return
	}

	// Ban all mentioned users
	for _, member := range mentionedMembers {
		if member.ID == author {
			session.ChannelMessageSend(channel, "<@"+member.ID+"> you cannot ban yourself.")
		} else {
			session.ChannelMessageSend(channel, "<@"+member.ID+"> was banned.")
			session.GuildBanCreate(message.GuildID, member.ID, 0)
		}
	}
}
