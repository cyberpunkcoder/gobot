package command

import (
	"github.com/bwmarrin/discordgo"
)

type help struct {
	command
}

func init() {
	help := help{command{
		name:        "help",
		parameters:  "(optional: @member)",
		description: "Lists commands avaliable to you or mentioned member.",
	}}
	executables = append(executables, &help)
}

func (h *help) execute(message *discordgo.MessageCreate, session *discordgo.Session) {
	author := message.Author.ID
	channel := message.ChannelID
	mentionedMembers := message.Mentions

	if len(mentionedMembers) == 0 {
		output := "**Commands Avaliable to You**\n"
		output += ">>> "
		output += generateHelpFor(author, channel, session)
		session.ChannelMessageSend(channel, output)
		return
	}
	// Send help to each mentioned member
	for _, member := range mentionedMembers {
		output := "**Commands Avaliable to <@" + member.ID + ">**\n"
		output += ">>> "
		output += generateHelpFor(member.ID, channel, session)
		session.ChannelMessageSend(channel, output)
	}
}

func generateHelpFor(memberID string, channelID string, session *discordgo.Session) string {
	output := ""

	userPermissions, _ := session.State.UserChannelPermissions(memberID, channelID)

	for _, executable := range executables {
		avaliable := true
		for _, permission := range executable.getPermisssions() {
			if userPermissions&permission == 0 {
				avaliable = false
				break
			}
		}

		if avaliable {
			output += "**" + executable.getName() + "** - *" + executable.getDescription() + "*\n"
			output += executable.getUsage() + "\n"
		}
	}
	return output
}
