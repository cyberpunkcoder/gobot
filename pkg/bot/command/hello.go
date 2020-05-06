package command

import "github.com/bwmarrin/discordgo"

// Hello messages "Hello!" in the chat
func Hello(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, "Hello!")
}
