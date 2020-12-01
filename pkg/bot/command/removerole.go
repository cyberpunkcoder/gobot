package command

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/reactionrole"
)

type removeRole struct {
	command
}

func init() {
	removeRole := removeRole{command{
		name:        "removerole",
		parameters:  "@role",
		description: "Removes a reaction role.",
		permissions: []int{discordgo.PermissionManageRoles},
	}}
	executables = append(executables, &removeRole)
}

func (r *removeRole) execute(message *discordgo.MessageCreate, session *discordgo.Session) {

	author := message.Author.ID
	commandChannel := message.ChannelID

	roleRegex := regexp.MustCompile(`<@&(\d+)>`)
	role := roleRegex.FindStringSubmatch(message.Content)

	if len(role) < 2 {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> "+r.wrongFormat())
		return
	}

	err := reactionrole.RemoveRole(role[1])

	// Return if unable to create message
	if err != nil {
		session.ChannelMessageSend(commandChannel, "<@"+author+"> role not found.")
		return
	}

	session.ChannelMessageSend(commandChannel, "<@"+author+"> role removed.")
}
