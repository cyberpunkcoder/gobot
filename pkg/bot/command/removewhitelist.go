package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/whitelist"
)

type removeWhitelist struct {
	command
}

func init() {
	removeWhitelist := removeWhitelist{command{
		name:        "removewhitelist",
		parameters:  "word or phrase",
		description: "Removes word or phrase whitelist for bot to ignore.",
		permissions: []int{discordgo.PermissionVoiceMuteMembers},
	}}
	executables = append(executables, &removeWhitelist)
}

func (a *removeWhitelist) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	author := message.Author.ID
	channel := message.ChannelID

	text := strings.Split(message.Content, " ")[0] + " "
	text = strings.TrimPrefix(message.Content, text)
	text = strings.ToLower(text)

	err := whitelist.RemoveWhitelist(text)

	// Return if whitelist not found
	if err != nil {
		session.ChannelMessageSend(channel, "<@"+author+"> whitelist not found.")
		return
	}

	session.ChannelMessageSend(channel, "<@"+author+"> whitelist removed.")
}
