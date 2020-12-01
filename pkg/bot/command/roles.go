package command

import (
	"log"

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

	// Remove command messaage
	session.ChannelMessageDelete(channel, message.ID)

	for _, catagory := range reactionrole.Catagories {

		output := "**" + catagory.Name + "**\n"
		output += "*React to give yourself a role.*\n"
		output += ">>> "

		for _, role := range catagory.Role {

			emoji, _ := session.State.Emoji(message.GuildID, role.Emoji.ID)
			output += "<:" + emoji.APIName() + "> - <@&" + role.ID + ">\n"
		}

		msg, err := session.ChannelMessageSend(message.ChannelID, output)

		// Return if unable to create message
		if err != nil {
			log.Println(err)
			return
		}

		reactionrole.SaveMessage(msg.ID)

		for _, role := range catagory.Role {

			// Save ID of the reaction role to reactionrolemessages.json to check for reactions later

			emoji, _ := session.State.Emoji(message.GuildID, role.Emoji.ID)
			err := session.MessageReactionAdd(channel, msg.ID, emoji.APIName())

			// Return if unable to create reaction
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
