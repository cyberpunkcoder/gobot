package command

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/role"
)

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
