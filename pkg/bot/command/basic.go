package command

import (
	"github.com/bwmarrin/discordgo"
)

// Hello messages "Hello!" in the chat
func Hello(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, "<@"+message.Author.ID+"> hello!")
}

// Unknown command was called
func Unknown(session *discordgo.Session, message *discordgo.MessageCreate, keyword string) {
	session.ChannelMessageSend(message.ChannelID, "<@"+message.Author.ID+"> unknown command \""+keyword+"\".")
}
