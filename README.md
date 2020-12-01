<img src="assets/discordgo.png" align="right" />

# Gobot
> A robust expandable discord bot written in go!

## Getting Started
Assuming you have go installed and have a discord bot with a token, simply copy and paste the command below!
You will be prompted to enter a discord bot token, command prefix and optionally the ID of a role users get when they join.
That's it! Assuming everything went well, your bot should be online and ready to use!

```git clone https://github.com/cyberpunkcoder/gobot.git && cd gobot/cmd/gobot && go build && ./gobot```
## Included Commands
> The following examples assume you have set the command prefix to be "!"
- <b>!help</b> or <b>!help @user</b> Bot will respond with a list of commands avaliable for you or mentioned user.
- <b>!hello</b> Pings the bot, bot will reply hello back.
- <b>!kick @user</b> Kicks a mentioned user from the guild.
- <b>!ban @user</b> Bans a mentioned user from the guild forever.
- <b>!purge</b> or <b>!purge @user</b> Removes last 100 messages from all users or mentioned user in channel.
- <b>!roles</b> Creates a reaction role selection menu with each reaction role in a catagory and associated emoji.
- <b>!addrole @role :emoji: Catagory </b> Creates a reaction role with an emoji and catagory.
- <b>!removerole @role</b> Removes a role from the reaction role menu.

## Credit
[Official Golang Website](https://golang.org/ "golang.org") | 
[Official Discordgo Repo](https://github.com/bwmarrin/discordgo "github.com/bwmarrin/discordgo")
