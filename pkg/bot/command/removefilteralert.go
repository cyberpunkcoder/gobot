package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/filter"
)

type removeFilterAlert struct {
	command
}

func init() {
	removeFilterAlert := removeFilterAlert{command{
		name:        "removefilteralert",
		parameters:  "(optional: @member)",
		description: "Removes alert for you or mentioned user for filter violations.",
		permissions: []int{discordgo.PermissionVoiceMuteMembers},
	}}
	executables = append(executables, &removeFilterAlert)
}

func (a *removeFilterAlert) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	author := message.Author.ID
	channel := message.ChannelID
	mentionedMembers := message.Mentions

	if len(mentionedMembers) > 0 {
		for _, member := range mentionedMembers {
			err := filter.RemoveAlert(*member)

			if err != nil {
				session.ChannelMessageSend(channel, "<@"+author+"> alert for member <@"+member.ID+"> does not exist.")
			} else {
				session.ChannelMessageSend(channel, "<@"+author+"> member <@"+member.ID+"> will no longer be alerted of filter violations.")
			}
		}
		return
	}

	err := filter.RemoveAlert(*message.Author)

	if err != nil {
		session.ChannelMessageSend(channel, "<@"+author+"> you were not set to be alerted of filter violations.")
	} else {
		session.ChannelMessageSend(channel, "<@"+author+"> you will no longer be alerted of filter violations.")
	}
}
