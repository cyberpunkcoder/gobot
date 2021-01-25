package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/whitelist"
)

type addWhitelist struct {
	command
}

func init() {
	addWhitelist := addWhitelist{command{
		name:        "addwhitelist",
		parameters:  "word or phrase",
		description: "Adds word or phrase whitelist for bot to ignore.",
		permissions: []int{discordgo.PermissionManageServer},
	}}
	executables = append(executables, &addWhitelist)
}

func (a *addWhitelist) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	author := message.Author.ID
	channel := message.ChannelID

	text := strings.Split(message.Content, " ")[0] + " "
	text = strings.TrimPrefix(message.Content, text)
	text = strings.ToLower(text)

	if strings.Contains(message.Content, " ") && text != "" {
		// Get keyword fom message text by removing everything after first space
		keyword := strings.Split(text, " ")[0]

		for _, e := range executables {
			// Check if there is a conflicting command
			if config.CommandPrefix+e.getName() == keyword {
				session.ChannelMessageSend(channel, "<@"+author+"> whitelist not added.")
				session.ChannelMessageSend(channel, "<@"+author+"> conflicts with existing command *"+config.CommandPrefix+e.getName()+"*.")
				return
			}
		}

		err := whitelist.SaveWhitelist(text)

		if err != nil {
			// Return if whitelist already exists
			session.ChannelMessageSend(channel, "<@"+author+"> whitelist already exists.")
			return
		}

		session.ChannelMessageSend(channel, "<@"+author+"> whitelist added.")
		return
	}
	session.ChannelMessageSend(channel, "<@"+author+"> "+a.wrongFormat())
}
