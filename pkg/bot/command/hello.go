package command

import (
	"github.com/bwmarrin/discordgo"
)

type hello struct {
	command
}

func init() {
	hello := hello{command{
		name:        "hello",
		description: "Replies hello to a user.",
	}}
	executables = append(executables, &hello)
}

func (h *hello) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	session.ChannelMessageSend(message.ChannelID, "<@"+message.Author.ID+"> hello!")
}
