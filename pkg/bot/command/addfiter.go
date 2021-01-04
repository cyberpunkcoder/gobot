package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/filter"
)

type addFilter struct {
	command
}

func init() {
	addFilter := addFilter{command{
		name:        "addfilter",
		parameters:  "word or phrase",
		description: "Adds word or phrase mute filter.",
		permissions: []int{discordgo.PermissionVoiceMuteMembers},
	}}
	executables = append(executables, &addFilter)
}

func (a *addFilter) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	author := message.Author.ID
	channel := message.ChannelID

	text := strings.Split(message.Content, " ")[0] + " "
	text = strings.TrimPrefix(message.Content, text)
	text = strings.ToLower(text)

	if strings.Contains(message.Content, " ") && text != "" {
		err := filter.SaveFilter(text)

		if err != nil {
			// Return if filter already exists
			session.ChannelMessageSend(channel, "<@"+author+"> filter already exists.")
			return
		}

		session.ChannelMessageSend(channel, "<@"+author+"> filter added.")
		return
	}
	session.ChannelMessageSend(channel, "<@"+author+"> "+a.wrongFormat())
}
