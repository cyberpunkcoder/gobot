package command

import (
	"github.com/bwmarrin/discordgo"
)

type addFilter struct {
	command
}

func init() {
	addRole := addRole{command{
		name:        "addfilter",
		parameters:  "word or phrase",
		description: "Adds word or phrase filter to mute members.",
		permissions: []int{discordgo.PermissionVoiceMuteMembers},
	}}
	executables = append(executables, &addRole)
}

func (a *addFilter) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	//author := message.Author.ID
	//channel := message.ChannelID

}
