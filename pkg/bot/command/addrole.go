package command

import (
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/reactionrole"
)

type addRole struct {
	command
}

func init() {
	addRole := addRole{command{
		name:        "addrole",
		parameters:  "@role :emoji: (optional: catagory)",
		description: "Adds a reaction role.",
		permissions: []int{discordgo.PermissionAddReactions, discordgo.PermissionManageRoles},
	}}
	executables = append(executables, &addRole)
}

func (a *addRole) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	author := message.Author.ID
	channel := message.ChannelID

	roleRegex := regexp.MustCompile(`<@&(\d+)>`)
	role := roleRegex.FindStringSubmatch(message.Content)

	emojiRegex := regexp.MustCompile(`<(a?):(.+):(\d+)>`)
	emoji := emojiRegex.FindStringSubmatch(message.Content)

	if len(role) != 2 || len(emoji) != 4 {
		session.ChannelMessageSend(channel, "<@"+author+"> "+a.wrongFormat())
		return
	}

	catagory := strings.ReplaceAll(message.Content, emoji[0], "")
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

	session.ChannelMessageSend(message.ChannelID, output)
}
