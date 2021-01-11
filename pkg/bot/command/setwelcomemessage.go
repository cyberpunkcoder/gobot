package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
)

type setWelcomeMessage struct {
	command
}

func init() {
	setWelcomeMessage := setWelcomeMessage{command{
		name:        "setwelcomemessage",
		parameters:  "welcome message",
		description: "Sets a welcome message bot sends members when they join.",
		permissions: []int{discordgo.PermissionManageServer},
	}}
	executables = append(executables, &setWelcomeMessage)
}

func (a *setWelcomeMessage) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	author := message.Author.ID
	channel := message.ChannelID

	if config.WelcomeMessage != "" {
		session.ChannelMessageSend(channel, "<@"+author+"> welcome message already set.")
		return
	}
	text := strings.Split(message.Content, " ")[0] + " "
	text = strings.TrimPrefix(message.Content, text)

	if strings.Contains(message.Content, " ") && text != "" {
		err := config.SaveWelcomeMessage(text)

		if err != nil {
			// Return if could not save welcome message
			session.ChannelMessageSend(channel, "<@"+author+"> could not set welcome message.")
			return
		}

		session.ChannelMessageSend(channel, "<@"+author+"> welcome message set.")
		return
	}
	session.ChannelMessageSend(channel, "<@"+author+"> "+a.wrongFormat())
}
