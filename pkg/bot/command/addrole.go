package command

import (
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

	emojiRegex := regexp.MustCompile(`[\x{00A9}-\x{1F6FF}|[\x{2600}-\x{26FF}]|\x{200D}`)
	emoji := emojiRegex.FindAllStringSubmatch(message.Content, -1)

	customEmojiRegex := regexp.MustCompile(`<(a?):(.+):(\d+)>`)
	customEmoji := customEmojiRegex.FindStringSubmatch(message.Content)

	if len(role) == 2 {
		catagory := message.Content
		newEmoji := reactionrole.Emoji{}

		if len(emoji) > 0 {
			// Loop through slice for a zwj combo emoji
			for _, i := range emoji {
				for j := 0; j < len(i); j++ {
					catagory = strings.ReplaceAll(catagory, i[j], "")
					newEmoji.Name += i[j]
				}
			}
		} else if len(customEmoji) == 4 {
			catagory = strings.ReplaceAll(catagory, customEmoji[0], "")
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
