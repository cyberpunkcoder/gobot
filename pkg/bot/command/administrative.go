package command

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/role"
)

// Kick all users mentioned in commandMessage from guild
func Kick(session *discordgo.Session, commandMessage *discordgo.MessageCreate) {

	author := commandMessage.Author.ID
	mentionedMembers := commandMessage.Mentions
	commandChannel := commandMessage.ChannelID

	// Bot permissions
	botPermissions, err := session.State.UserChannelPermissions(session.State.User.ID, commandChannel)

	// Return if unable to check bot permissions
	if err != nil {
		log.Println(err)
		session.ChannelMessageSend(commandChannel, "<@"+author+"> unable to check my permissions.")
		return
	}

	// Return if bot does not have permission to kick
	if botPermissions&discordgo.PermissionKickMembers == 0 {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> I don't have permission to kick.")
		return
	}

	// Command author permissions
	permissions, err := session.State.UserChannelPermissions(author, commandChannel)

	// Return if unable to check command author permissions
	if err != nil {
		log.Println(err)
		session.ChannelMessageSend(commandChannel, "<@"+author+"> unable to check your permissions.")
		return
	}

	// Return if author does not have permission to kick
	if permissions&discordgo.PermissionKickMembers == 0 {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> you don't have permission to kick.")
		return
	}

	// Return if no users mentioned in command
	if len(mentionedMembers) <= 0 {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> you must mention a user to kick.")
		return
	}

	// Kick all mentioned users
	for _, member := range mentionedMembers {
		if member.ID == author {
			session.ChannelMessageSend(commandChannel, "<@"+member.ID+"> you cannot kick yourself.")
		} else {
			session.ChannelMessageSend(commandChannel, "<@"+member.ID+"> was kicked.")
			session.GuildMemberDelete(commandMessage.GuildID, member.ID)
		}
	}
}

// Ban all users mentioned in commandMessage from guild
func Ban(session *discordgo.Session, commandMessage *discordgo.MessageCreate) {

	author := commandMessage.Author.ID
	mentionedMembers := commandMessage.Mentions
	commandChannel := commandMessage.ChannelID

	// Bot permissions
	botPermissions, err := session.State.UserChannelPermissions(session.State.User.ID, commandChannel)

	// Return if unable to check bot permissions
	if err != nil {
		log.Println(err)
		session.ChannelMessageSend(commandChannel, "<@"+author+"> unable to check my permissions.")
		return
	}

	// Return if bot does not have permission to ban
	if botPermissions&discordgo.PermissionBanMembers == 0 {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> I don't have permission to ban.")
		return
	}

	// Command author permissions
	permissions, err := session.State.UserChannelPermissions(author, commandChannel)

	// Return if unable to check command author permissions
	if err != nil {
		log.Println(err)
		session.ChannelMessageSend(commandChannel, "<@"+author+"> unable to check your permissions.")
		return
	}

	// Return if command author does not have permission to ban
	if permissions&discordgo.PermissionBanMembers == 0 {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> you don't have permission to ban.")
		return
	}

	// Return if no users mentioned in command
	if len(mentionedMembers) <= 0 {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> you must mention a user to ban.")
		return
	}

	// Ban all mentioned users
	for _, member := range mentionedMembers {
		if member.ID == author {
			session.ChannelMessageSend(commandChannel, "<@"+member.ID+"> you cannot ban yourself.")
		} else {
			session.ChannelMessageSend(commandChannel, "<@"+member.ID+"> was banned.")
			session.GuildBanCreate(commandMessage.GuildID, member.ID, 0)
		}
	}
}

// Purge (remove) last 100 messages from the chat or any mentioned user
func Purge(session *discordgo.Session, commandMessage *discordgo.MessageCreate) {

	author := commandMessage.Author.ID
	mentionedMembers := commandMessage.Mentions
	commandChannel := commandMessage.ChannelID

	// Bot permissions
	botPermissions, err := session.State.UserChannelPermissions(session.State.User.ID, commandChannel)

	// Last 100 messages in command channel
	messages, err := session.ChannelMessages(commandChannel, 100, commandMessage.ID, "", "")

	// Return if unable to check bot permissions
	if err != nil {
		log.Println(err)
		session.ChannelMessageSend(commandChannel, "<@"+author+"> unable to check my permissions.")
		return
	}

	// Return if bot does not have permission to delete messages
	if botPermissions&discordgo.AuditLogActionMessageDelete == 0 {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> I don't have permission to delete messages.")
		return
	}

	// Author permissions
	permissions, err := session.State.UserChannelPermissions(author, commandChannel)

	// Return if unable to check command author permissions
	if err != nil {
		log.Println(err)
		session.ChannelMessageSend(commandChannel, "<@"+author+"> unable to check your permissions.")
		return
	}

	// Return if command author does not have permission to delete messages
	if permissions&discordgo.AuditLogActionMessageDelete == 0 {
		log.Println(err)
		session.ChannelMessageSend(commandChannel, "<@"+author+"> you do not have permission to delete messages.")
		return
	}

	// Return if unable to check messages
	if err != nil {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> unable to check messages.")
		return
	}

	session.ChannelMessageSend(commandChannel, "<@"+author+"> purging.")

	// For each of the last 100 messages in the channel
	for _, message := range messages {

		// If no user was mentioned in the command delete each of last 100 messages
		if len(mentionedMembers) <= 0 {
			session.ChannelMessageDelete(commandChannel, message.ID)
		} else {

			// For each of the users mentioned in the command
			for _, member := range mentionedMembers {
				if message.Author.ID == member.ID {
					session.ChannelMessageDelete(commandChannel, message.ID)
				}
			}
		}
	}
}

// Roles creates a menu for users to choose roles by reaction
func Roles(session *discordgo.Session, commandMessage *discordgo.MessageCreate) {

	author := commandMessage.Author.ID
	commandChannel := commandMessage.ChannelID

	// Remove command messaage
	session.ChannelMessageDelete(commandChannel, commandMessage.ID)

	// Bot permissions
	botPermissions, err := session.State.UserChannelPermissions(session.State.User.ID, commandChannel)

	// Return if unable to check bot permissions
	if err != nil {
		log.Println(err)
		session.ChannelMessageSend(commandChannel, "<@"+author+"> unable to check my permissions.")
		return
	}

	// Return if bot does not have permission to add reaction
	if botPermissions&discordgo.PermissionAddReactions == 0 {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> I don't have permission to add reactions.")
		return
	}

	// Return if bot does not have permission to manage roles
	if botPermissions&discordgo.PermissionManageRoles == 0 {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> I don't have permission to manage roles.")
		return
	}

	// Command author permissions
	permissions, err := session.State.UserChannelPermissions(author, commandChannel)

	// Return if unable to check author permissions
	if err != nil {
		log.Println(err)
		session.ChannelMessageSend(commandChannel, "<@"+author+"> unable to check your permissions.")
		return
	}

	// Return if author does not have permission to add reactions
	if permissions&discordgo.PermissionAddReactions == 0 {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> you don't have permission to add reactions.")
		return
	}

	// Return if author does not have permission to manage roles
	if permissions&discordgo.PermissionManageRoles == 0 {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> you don't have permission to manage roles.")
		return
	}

	for _, catagory := range role.ReactionRoleCatagories {

		output := "**" + catagory.Name + "**\n"
		output += "*React to give yourself a role.*\n"
		output += ">>> "

		for _, role := range catagory.Role {
			emoji, _ := session.State.Emoji(commandMessage.GuildID, role.EmojiID)
			output += "<:" + emoji.APIName() + "> - <@&" + role.RoleID + ">\n"
		}

		msg, err := session.ChannelMessageSend(commandMessage.ChannelID, output)

		// Return if unable to create message
		if err != nil {
			log.Println(err)
			session.ChannelMessageSend(commandChannel, "<@"+author+"> unable to create message.")
			return
		}

		// Save ID of the reaction role to reactionrolemessages.json to check for reactions later
		role.SaveReactionRoleMessage(msg.ID)

		for _, role := range catagory.Role {

			emoji, _ := session.State.Emoji(commandMessage.GuildID, role.EmojiID)
			err := session.MessageReactionAdd(commandChannel, msg.ID, emoji.APIName())

			// Return if unable to create reaction
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
