package command

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Kick all users mentioned in message from guild
func Kick(session *discordgo.Session, message *discordgo.MessageCreate) {

	// Bot permissions
	botPermissions, err := session.State.UserChannelPermissions(session.State.User.ID, message.ChannelID)

	// Command author and permissions
	author := message.Author.ID

	if err != nil {
		fmt.Println(err)
		session.ChannelMessageSend(message.ChannelID, "<@"+author+"> unable to check my permissions.")
	} else if botPermissions&discordgo.PermissionKickMembers == 0 {
		session.ChannelMessageSend(message.ChannelID, "<@"+author+"> I don't have permission to kick.")
	} else {

		// Author permissions
		permissions, err := session.State.UserChannelPermissions(author, message.ChannelID)

		if err != nil {
			fmt.Println(err)
			session.ChannelMessageSend(message.ChannelID, "<@"+author+"> unable to check your permissions.")
		} else if permissions&discordgo.PermissionKickMembers == 0 {
			session.ChannelMessageSend(message.ChannelID, "<@"+author+"> you don't have permission to kick.")
		} else {
			members := message.Mentions // All users mentioned in command

			if len(members) <= 0 {
				session.ChannelMessageSend(message.ChannelID, "<@"+author+"> you must mention a user to kick.")
			} else {

				// Kick all mentioned users
				for _, member := range members {
					session.ChannelMessageSend(message.ChannelID, "<@"+member.ID+"> was kicked.")
					session.GuildMemberDelete(message.GuildID, member.ID)
				}
			}
		}

	}
}

// Ban all users mentioned in message from guild
func Ban(session *discordgo.Session, message *discordgo.MessageCreate) {

	// Bot permissions
	botPermissions, err := session.State.UserChannelPermissions(session.State.User.ID, message.ChannelID)

	// Command author and permissions
	author := message.Author.ID

	if err != nil {
		fmt.Println(err)
		session.ChannelMessageSend(message.ChannelID, "<@"+author+"> unable to check my permissions.")
	} else if botPermissions&discordgo.PermissionBanMembers == 0 {
		session.ChannelMessageSend(message.ChannelID, "<@"+author+"> I don't have permission to ban.")
	} else {

		// Author permissions
		permissions, err := session.State.UserChannelPermissions(author, message.ChannelID)

		if err != nil {
			fmt.Println(err)
			session.ChannelMessageSend(message.ChannelID, "<@"+author+"> unable to check your permissions.")
		} else if permissions&discordgo.PermissionBanMembers == 0 {
			session.ChannelMessageSend(message.ChannelID, "<@"+author+"> you don't have permission to ban.")
		} else {
			members := message.Mentions // All users mentioned in command

			if len(members) <= 0 {
				session.ChannelMessageSend(message.ChannelID, "<@"+author+"> you must mention a user to ban.")
			} else {

				// Ban all mentioned users
				for _, member := range members {
					session.ChannelMessageSend(message.ChannelID, "<@"+member.ID+"> was banned.")
					session.GuildBanCreate(message.GuildID, member.ID, 0)
				}
			}
		}
	}
}
