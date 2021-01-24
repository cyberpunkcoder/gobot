package bot

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/command"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/filter"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/reactionrole"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/whitelist"
)

// ready is called whenever the bot has successfully logged in
func ready(session *discordgo.Session, ready *discordgo.Ready) {

	// Display name of bot to user
	log.Println("Logged in as", session.State.User, "press CTRL-C to exit")
}

// guildMemberAdd is called when a member joins the guild
func guildMemberAdd(session *discordgo.Session, user *discordgo.GuildMemberAdd) {
	if config.JoinRole != "" {
		// Give a member the join role
		session.GuildMemberRoleAdd(user.GuildID, user.User.ID, config.JoinRole)
	}

	if config.WelcomeMessage != "" {
		// Send welcome message to new member
		welcomeChannel, _ := session.UserChannelCreate(user.User.ID)
		session.ChannelMessageSend(welcomeChannel.ID, "<@"+user.User.ID+">, "+config.WelcomeMessage)
	}
}

// messageCreate is called whenever a message has been created
func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ensure that message was not created by bot itself
	if message.Author.ID != session.State.User.ID {
		// Check if message is a bot command
		if strings.HasPrefix(message.Content, config.CommandPrefix) {
			// Check if message contains a whitelist
			if whitelist.Check(message.Content) {
				command.Execute(message, session)
			}
		}
		// Check if message violates filter
		filter.Check(message, session)
	}
}

// messageReactionAdd is called whenever a reaction has been added to a message
func messageReactionAdd(session *discordgo.Session, reaction *discordgo.MessageReactionAdd) {

	// Bot permissions
	botPermissions, err := session.State.UserChannelPermissions(session.State.User.ID, reaction.ChannelID)

	// Return if unable to check bot permissions
	if err != nil {
		log.Println(err)
		return
	}

	// Return if bot does not have permission to manage roles
	if botPermissions&discordgo.PermissionManageRoles == 0 {
		session.ChannelMessage(reaction.ChannelID, "<@"+reaction.UserID+"> I don't have permission to manage roles.")
		return
	}

	for _, reactionRoleMessage := range reactionrole.Messages {
		if reaction.MessageID == reactionRoleMessage {
			for _, reactionRoleCatagory := range reactionrole.Catagories {
				for _, reactionRole := range reactionRoleCatagory.Role {
					if reactionRole.Emoji.ID == reaction.Emoji.ID && reaction.UserID != session.State.User.ID {
						session.GuildMemberRoleAdd(reaction.GuildID, reaction.UserID, reactionRole.ID)
					}
				}
			}
		}
	}
}

// messageReactionAdd is called whenever a reaction has been removed from a message
func messageReactionRemove(session *discordgo.Session, reaction *discordgo.MessageReactionRemove) {

	// Bot permissions
	botPermissions, err := session.State.UserChannelPermissions(session.State.User.ID, reaction.ChannelID)

	// Return if unable to check bot permissions
	if err != nil {
		log.Println(err)
		return
	}

	// Return if bot does not have permission to manage roles
	if botPermissions&discordgo.PermissionManageRoles == 0 {
		session.ChannelMessage(reaction.ChannelID, "<@"+reaction.UserID+"> I don't have permission to manage roles.")
		return
	}

	for _, reactionRoleMessage := range reactionrole.Messages {
		if reaction.MessageID == reactionRoleMessage {
			for _, reactionRoleCatagory := range reactionrole.Catagories {
				for _, reactionRole := range reactionRoleCatagory.Role {
					if reactionRole.Emoji.ID == reaction.Emoji.ID {
						session.GuildMemberRoleRemove(reaction.GuildID, reaction.UserID, reactionRole.ID)
					}
				}
			}
		}
	}
}
