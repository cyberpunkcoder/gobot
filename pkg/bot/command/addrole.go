package command

import (
	"fmt"
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

	fmt.Println(message.Content)

	author := message.Author.ID
	channel := message.ChannelID

	roleRegex := regexp.MustCompile(`<@&(\d+)>`)
	role := roleRegex.FindStringSubmatch(message.Content)

	emojiRegex := regexp.MustCompile(`[\x{1F600}-\x{1F6FF}]|[\x{2600}-\x{26FF}]`)
	emoji := emojiRegex.FindStringSubmatch(message.Content)

	regionalRegex := regexp.MustCompile(`[\x{1F1E6}-\x{1F1FF}]`)
	region := regionalRegex.FindStringSubmatch(message.Content)

	for _, e := range region {
		fmt.Println(e)
	}

	customEmojiRegex := regexp.MustCompile(`<(a?):(.+):(\d+)>`)
	customEmoji := customEmojiRegex.FindStringSubmatch(message.Content)

	if len(role) == 2 {
		catagory := message.Content

		if len(emoji) == 1 {
			catagory = strings.ReplaceAll(message.Content, emoji[0], "")
		} else if len(customEmoji) == 4 {
			catagory = strings.ReplaceAll(message.Content, customEmoji[0], "")
		} else if len(region) == 1 || len(region) == 2 {
			catagory = strings.ReplaceAll(message.Content, region[0], "")
		}

		catagory = strings.TrimLeft(catagory, " ")
		catagory = strings.ReplaceAll(catagory, role[0], "")

		spaceRegex := regexp.MustCompile(`\s(.*)`)
		catagory = spaceRegex.FindString(catagory)
		catagory = strings.TrimSpace(catagory)

		newRole := reactionrole.Role{
			ID: role[1],
			Emoji: reactionrole.Emoji{
				Prefix: customEmoji[1],
				Name:   customEmoji[2],
				ID:     customEmoji[3],
			},
		}

		reactionrole.SaveRole(catagory, newRole)

		output := "<@" + author + "> role added.\n"
		output += ">>> "
		output += "<" + customEmoji[1] + ":" + customEmoji[2] + ":" + customEmoji[3] + "> - <@&" + role[1] + ">\n"

		session.ChannelMessageSend(message.ChannelID, output)
	}
	session.ChannelMessageSend(channel, "<@"+author+"> "+a.wrongFormat())
}
