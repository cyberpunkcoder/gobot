<img src="assets/discordgo.png" align="right" />

# Gobot
> A robust expandable discord bot written in go!

## Getting Started
Assuming you have go installed and have a discord bot with a token, simply copy and paste the command below!
You will be prompted via command line to enter a discord bot token and command prefix.
That's it! Assuming everything went well, your bot should be online and ready to use!

```git clone https://github.com/cyberpunkcoder/gobot.git && cd gobot/cmd/gobot && go build && ./gobot```
## Included Commands
> The following examples assume you have set the command prefix to be "!"
### Auto Roles
- <b>!setjoinrole @role</b> Sets the role users will be given when they join.
- <b>!removejoinrole</b> Removes the role users will be given when they join.
- <b>!setmuterole @role</b> Sets the role users will be given when they are muted.
- <b>!removemuterole</b> Removes the role users will be given when they are muted.

### User Tools
- <b>!help</b> or <b>!help @user</b> Bot will respond with a list of commands avaliable for you or mentioned user.
- <b>!hello</b> Pings the bot, bot will reply hello back.
- <b>!kick @user</b> Kicks a mentioned user from the guild.
- <b>!ban @user</b> Bans a mentioned user from the guild forever.
- <b>!purge</b> or <b>!purge @user</b> Removes last 100 messages from all users or mentioned user in channel.

### Reaction Roles
- <b>!roles</b> Creates a reaction role selection menu with each reaction role in a catagory and associated emoji.
- <b>!addrole @role :emoji: Catagory </b> Creates a reaction role with an emoji and catagory.
- <b>!removerole @role</b> Removes a role from the reaction role menu.

### Word or Phrase Filters
- <b>!addfilter word or phrase</b> Creates a message filter that if volated a message will be removed and user muted if mute role set.
- <b>!removefilter word or phrase</b> Removes a message filter.
- <b>!addfilteralert</b> or <b>!addfilteralert @user</b> Adds you or user to a list of people to be alerted if a filter is violated.
- <b>!removefilteralert</b> or <b>!removefilteralert @user</b> Removes you or user to a list of people to be alerted if a filter is violated.
- <b>!filters</b> Lists word or phrase filters.

## Credit
[Official Golang Website](https://golang.org/ "golang.org") | 
[Official Discordgo Repo](https://github.com/bwmarrin/discordgo "github.com/bwmarrin/discordgo")
