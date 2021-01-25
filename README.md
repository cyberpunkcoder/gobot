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
- ```!setjoinrole @role``` Sets the role members will be given when they join.
- ```!removejoinrole``` Removes the role members will be given when they join.
- ```!setmuterole @role``` Sets the role members will be given when they are muted.
- ```!removemuterole``` Removes the role members will be given when they are muted.

### Welcome Message
- ```!welcome``` or ```!welcome @member``` Bot will send welcome message in channel or addressed to member.
- ```!setwelcomemessage welcome message``` Sets the welcome message members will be sent in DM when they join.
- ```!removewelcomemessage``` Removes the welcome message.

### User Tools
- ```!help``` or ```!help @member``` Bot will respond with a list of commands avaliable for you or mentioned member.
- ```!hello``` Pings the bot, bot will reply hello back.
- ```!kick @member``` Kicks a mentioned member from the guild.
- ```!ban @member``` Bans a mentioned member from the guild forever.
- ```!purge``` or ```!purge @member``` Removes last 100 messages from all members or mentioned member in channel.

### Reaction Roles
- ```!roles``` Creates a reaction role selection menu with each reaction role in a catagory and associated emoji.
- ```!addrole @role :emoji: Catagory``` Creates a reaction role with an emoji and catagory.
- ```!removerole @role``` Removes a role from the reaction role menu.

### Whitelists
- ```!whitelists``` Lists all whitelisted commands for the bot to ignore.
- ```!addwhitelist word or phrase``` Adds a whitelisted word or phrase for bot to ignore
- ```!removewhitelist word or phrase``` Removes a whitelisted word or phrase for bot to ignore.

### Word or Phrase Filters
- ```!addfilter word or phrase``` Creates a message filter that if volated a message will be removed and member muted if mute role set.
- ```!removefilter word or phrase``` Removes a message filter.
- ```!addfilteralert``` or ```!addfilteralert @member``` Adds you or member to a list of people to be alerted if a filter is violated.
- ```!removefilteralert``` or ```!removefilteralert @member``` Removes you or member to a list of people to be alerted if a filter is violated.
- ```!filters``` Lists word or phrase filters.

## Credit
[Official Golang Website](https://golang.org/ "golang.org") | 
[Official Discordgo Repo](https://github.com/bwmarrin/discordgo "github.com/bwmarrin/discordgo")
