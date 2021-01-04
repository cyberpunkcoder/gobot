package command

import (
	"github.com/bwmarrin/discordgo"
)

type removeFilter struct {
	command
}

func init() {
	removeFilter := removeFilter{command{
		name:        "removefilter",
		parameters:  "word or phrase",
		description: "Removes a word or phrase mute filter.",
		permissions: []int{discordgo.PermissionVoiceMuteMembers},
	}}
	executables = append(executables, &removeFilter)
}

func (a *removeFilter) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	//author := message.Author.ID
	//channel := message.ChannelID

}
