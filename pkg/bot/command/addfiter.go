package command

import (
	"github.com/bwmarrin/discordgo"
)

type addFilter struct {
	command
}

func init() {
	addFilter := addFilter{command{
		name:        "addfilter",
		parameters:  "word or phrase",
		description: "Adds word or phrase mute filter.",
		permissions: []int{discordgo.PermissionVoiceMuteMembers},
	}}
	executables = append(executables, &addFilter)
}

func (a *addFilter) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	//author := message.Author.ID
	//channel := message.ChannelID

}
