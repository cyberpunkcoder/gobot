package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/filter"
)

type addFilterAlert struct {
	command
}

func init() {
	addFilterAlert := addFilterAlert{command{
		name:        "addfilteralert",
		parameters:  "(optional: @member)",
		description: "Adds alert for you or mentioned user for filter violations.",
		permissions: []int{discordgo.PermissionVoiceMuteMembers},
	}}
	executables = append(executables, &addFilterAlert)
}

func (a *addFilterAlert) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	author := message.Author.ID
	channel := message.ChannelID
	mentionedMembers := message.Mentions

	if len(mentionedMembers) > 0 {
		for _, member := range mentionedMembers {
			err := filter.SaveAlert(*member)

			if err != nil {
				session.ChannelMessageSend(channel, "<@"+author+"> alert for member <@"+member.ID+"> already exists.")
			} else {
				session.ChannelMessageSend(channel, "<@"+author+"> member <@"+member.ID+"> will be alerted of filter violations.")
			}
		}
		return
	}

	err := filter.SaveAlert(*message.Author)

	if err != nil {
		session.ChannelMessageSend(channel, "<@"+author+"> you are already set to be alerted of filter violations.")
	} else {
		session.ChannelMessageSend(channel, "<@"+author+"> you will be alerted of filter violations.")
	}
}
