package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/filter"
)

type filters struct {
	command
}

func init() {
	filters := filters{command{
		name:        "filters",
		description: "Lists mute filters.",
		permissions: []int{discordgo.PermissionVoiceMuteMembers},
	}}
	executables = append(executables, &filters)
}

func (a *filters) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	channel := message.ChannelID

	output := "**Mute Filters**\n"
	output += ">>> "

	for _, filter := range filter.Filters {
		output += "**" + filter.Text + "**\n"
	}

	session.ChannelMessageSend(channel, output)
}
