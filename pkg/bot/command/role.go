package command

import (
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/reactionrole"
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

	for _, catagory := range reactionrole.Catagories {

		output := "**" + catagory.Name + "**\n"
		output += "*React to give yourself a role.*\n"
		output += ">>> "

		for _, role := range catagory.Role {

			emoji, _ := session.State.Emoji(commandMessage.GuildID, role.Emoji.ID)
			output += "<:" + emoji.APIName() + "> - <@&" + role.ID + ">\n"
		}

		msg, err := session.ChannelMessageSend(commandMessage.ChannelID, output)

		// Return if unable to create message
		if err != nil {
			log.Println(err)
			session.ChannelMessageSend(commandChannel, "<@"+author+"> unable to create message.")
			return
		}

		reactionrole.SaveMessage(msg.ID)

		for _, role := range catagory.Role {

			// Save ID of the reaction role to reactionrolemessages.json to check for reactions later

			emoji, _ := session.State.Emoji(commandMessage.GuildID, role.Emoji.ID)
			err := session.MessageReactionAdd(commandChannel, msg.ID, emoji.APIName())

			// Return if unable to create reaction
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

// AddRole adds a role to the reaction role menu
func AddRole(session *discordgo.Session, commandMessage *discordgo.MessageCreate) {

	author := commandMessage.Author.ID

	roleRegex := regexp.MustCompile(`<@&(\d+)>`)
	role := roleRegex.FindStringSubmatch(commandMessage.Content)

	emojiRegex := regexp.MustCompile(`<(a?):(.+):(\d+)>`)
	emoji := emojiRegex.FindStringSubmatch(commandMessage.Content)

	catagory := strings.ReplaceAll(commandMessage.Content, emoji[0], "")
	catagory = strings.TrimLeft(catagory, " ")
	catagory = strings.ReplaceAll(catagory, role[0], "")

	spaceRegex := regexp.MustCompile(`\s(.*)`)
	catagory = spaceRegex.FindString(catagory)
	catagory = strings.TrimSpace(catagory)

	newRole := reactionrole.Role{
		ID: role[1],
		Emoji: reactionrole.Emoji{
			Prefix: emoji[1],
			Name:   emoji[2],
			ID:     emoji[3],
		},
	}

	reactionrole.SaveRole(catagory, newRole)

	output := "<@" + author + "> role added.\n"
	output += ">>> "
	output += "<" + emoji[1] + ":" + emoji[2] + ":" + emoji[3] + "> - <@&" + role[1] + ">\n"

	session.ChannelMessageSend(commandMessage.ChannelID, output)
}

// RemoveRole adds a role to the reaction role menu
func RemoveRole(session *discordgo.Session, commandMessage *discordgo.MessageCreate) {

	author := commandMessage.Author.ID
	commandChannel := commandMessage.ChannelID

	roleRegex := regexp.MustCompile(`<@&(\d+)>`)
	role := roleRegex.FindStringSubmatch(commandMessage.Content)

	err := reactionrole.RemoveRole(role[1])

	// Return if unable to create message
	if err != nil {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> role not found.")
		return
	}

	session.ChannelMessageSend(commandChannel, "<@"+author+"> role removed.")
}
