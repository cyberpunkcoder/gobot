package bot

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/command"
	"github.com/cyberpunkprogrammer/gobot/pkg/config"
)

// ready is called whenever the bot has successfully logged in
func ready(session *discordgo.Session, ready *discordgo.Ready) {

	// Display name of bot to user
	log.Println("Logged in as", session.State.User, "press CTRL-C to exit")
}

// guildMemberAdd is called when a member joins the guild
func guildMemberAdd(session *discordgo.Session, user *discordgo.GuildMemberAdd) {
	session.GuildMemberRoleAdd(user.GuildID, user.User.ID, config.JoinRole)
}

// messageCreate is called whenever a message has been created
func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	if strings.HasPrefix(message.Content, config.CommandPrefix) {

		// Get keyword fom message text by removing CommandPrefix and everything after first space
		keyword := strings.Replace(message.Content, config.CommandPrefix, "", -1)
		keyword = strings.Split(keyword, " ")[0]

		// Ignore all messages created by the bot itself
		if message.Author.ID == session.State.User.ID {
			return
		}

		switch keyword {
		case "hello":
			command.Hello(session, message)
		case "kick":
			command.Kick(session, message)
		case "ban":
			command.Ban(session, message)
		case "purge":
			command.Purge(session, message)
		default:
			command.Unknown(session, message, keyword)
		}
	}
}
