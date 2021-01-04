package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/filter"
)

type removeFilter struct {
	command
}

func init() {
	removeFilter := removeFilter{command{
		name:        "removefilter",
		parameters:  "word or phrase",
		description: "Removes word or phrase mute filter.",
		permissions: []int{discordgo.PermissionVoiceMuteMembers},
	}}
	executables = append(executables, &removeFilter)
}

func (a *removeFilter) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	author := message.Author.ID
	channel := message.ChannelID

	text := strings.Split(message.Content, " ")[0] + " "
	text = strings.TrimPrefix(message.Content, text)
	text = strings.ToLower(text)

	err := filter.RemoveFilter(text)

	// Return if filter not found
	if err != nil {
		session.ChannelMessageSend(channel, "<@"+author+"> filter not found.")
		return
	}

	session.ChannelMessageSend(channel, "<@"+author+"> filter removed.")
}
