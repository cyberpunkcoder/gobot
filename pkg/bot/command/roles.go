package command

import (
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/reactionrole"
)

type roles struct {
	command
}

func init() {
	roles := roles{command{
		name:        "roles",
		description: "Lists roles for members to react too.",
		permissions: []int{discordgo.PermissionAddReactions, discordgo.PermissionManageRoles},
	}}
	executables = append(executables, &roles)
}

func (r *roles) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	channel := message.ChannelID
	session.ChannelMessageDelete(channel, message.ID)
	checkingMessage, _ := session.ChannelMessageSend(message.ChannelID, "<@"+message.Author.ID+"> checking access to each emoji.")

	// Check if bot has access to all emojis
	for _, catagory := range reactionrole.Catagories {
		for _, role := range catagory.Role {
			if role.Emoji.ID != "" {

				err := session.MessageReactionAdd(checkingMessage.ChannelID, checkingMessage.ID, role.Emoji.APIName())

				// Return if unable to create reaction
				if err != nil {
					if strings.Contains(err.Error(), strconv.Itoa(discordgo.ErrCodeUnknownEmoji)) {
						session.ChannelMessageDelete(channel, message.ID)
						session.ChannelMessageDelete(channel, checkingMessage.ID)
						session.ChannelMessageSend(message.ChannelID, "<@"+message.Author.ID+"> I don't have access to the emoji "+role.Emoji.ToString()+" .")
						return
					}
					log.Println(err)
					return
				}
			}
		}
	}

	session.ChannelMessageDelete(channel, checkingMessage.ID)

	// Create the reaction role menu
	for _, catagory := range reactionrole.Catagories {
		output := ""

		if catagory.Name != "" {
			output = "**" + catagory.Name + "**\n"
		}

		output += "*React to give yourself a role.*\n"
		output += ">>> "

		for _, role := range catagory.Role {
			output += role.Emoji.ToString() + " - <@&" + role.ID + ">\n"
		}

		msg, err := session.ChannelMessageSend(message.ChannelID, output)

		// Return if unable to create message
		if err != nil {
			log.Println(err)
			return
		}

		// Save ID of the reaction role to check for reactions later
		reactionrole.SaveMessage(msg.ID)

		for _, role := range catagory.Role {
			err = session.MessageReactionAdd(channel, msg.ID, role.Emoji.APIName())

			// Return if unable to create reaction
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
