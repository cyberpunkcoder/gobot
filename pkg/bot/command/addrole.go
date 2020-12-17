package command

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
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

	emojiRegex := regexp.MustCompile(`[\x{1F600}-\x{1F6FF}]|[\x{2600}-\x{26FF}]`)
	emoji := emojiRegex.FindStringSubmatch(message.Content)

	customEmojiRegex := regexp.MustCompile(`<(a?):(.+):(\d+)>`)
	customEmoji := customEmojiRegex.FindStringSubmatch(message.Content)

	regionalRegex := regexp.MustCompile(`[\x{1F1E6}-\x{1F1FF}]`)
	region := regionalRegex.FindAllStringSubmatch(message.Content, -1)

	if len(role) == 2 {
		catagory := message.Content
		newEmoji := reactionrole.Emoji{}

		if len(emoji) == 1 {
			catagory = strings.ReplaceAll(message.Content, emoji[0], "")
			newEmoji.Name = emoji[0]

		} else if len(customEmoji) == 4 {
			catagory = strings.ReplaceAll(message.Content, customEmoji[0], "")
			newEmoji.Prefix = customEmoji[1]
			newEmoji.Name = customEmoji[2]
			newEmoji.ID = customEmoji[3]

			// Check if bot has access to custom emoji
			emojiText := newEmoji.Name + ":" + newEmoji.ID
			checkingMessage, _ := session.ChannelMessageSend(message.ChannelID, "<@"+message.Author.ID+"> checking access to emoji.")
			err := session.MessageReactionAdd(checkingMessage.ChannelID, checkingMessage.ID, emojiText)

			if err != nil {
				if strings.Contains(err.Error(), strconv.Itoa(discordgo.ErrCodeUnknownEmoji)) {
					session.ChannelMessageDelete(channel, checkingMessage.ID)
					session.ChannelMessageSend(message.ChannelID, "<@"+message.Author.ID+"> I don't have access to the emoji "+newEmoji.ToString()+" .")
					return
				}
				log.Println(err)
				return
			}

			// Remove checking message
			session.ChannelMessageDelete(checkingMessage.ChannelID, checkingMessage.ID)

		} else if len(region) > 1 {
			if len(region) == 1 {
				catagory = strings.ReplaceAll(message.Content, region[0][0], "")
				newEmoji.Name = region[0][0]
			} else if len(region) == 2 {
				name := region[0][0] + region[1][0]
				catagory = strings.ReplaceAll(message.Content, name, "")
				newEmoji.Name = name
			}
		} else {
			// Wrong emoji format
			session.ChannelMessageSend(channel, "<@"+author+"> "+a.wrongFormat())
			return
		}

		catagory = strings.TrimLeft(catagory, " ")
		catagory = strings.ReplaceAll(catagory, role[0], "")
		spaceRegex := regexp.MustCompile(`\s(.*)`)
		catagory = spaceRegex.FindString(catagory)
		catagory = strings.TrimSpace(catagory)

		fmt.Println(catagory)

		newRole := reactionrole.Role{
			ID:    role[1],
			Emoji: newEmoji,
		}

		// Save role to disk
		reactionrole.SaveRole(catagory, newRole)

		output := "<@" + author + "> role added.\n"
		output += ">>> "
		output += newEmoji.ToString() + " - <@&" + role[1] + ">\n"

		session.ChannelMessageSend(message.ChannelID, output)
		return
	}
	// Wrong role format
	session.ChannelMessageSend(channel, "<@"+author+"> "+a.wrongFormat())
}
