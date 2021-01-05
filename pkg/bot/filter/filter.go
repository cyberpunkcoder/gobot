package filter

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cyberpunkprogrammer/gobot/pkg/bot/config"
)

var (
	// BinPath is the location of the executable binary
	BinPath string

	// Filters that members will be muted for messaging
	Filters []filter

	// Alerts members if a filter is violated
	Alerts []alert

	filtersFile = "/json/filters.json"
	alertsFile  = "/json/filteralerts.json"
)

// Filter phrase or word members will be muted for messaging
type filter struct {
	Text string `json:"text"`
}

//alert member that will be alerted if filtr is violated
type alert struct {
	ID string `json:"ID"`
}

// Check a message against the filter
func Check(message *discordgo.MessageCreate, session *discordgo.Session) {
	var violations []string
	for _, filter := range Filters {
		if strings.Contains(message.Content, filter.Text) && !message.Author.Bot {
			violations = append(violations, filter.Text)
		}
	}

	if len(violations) > 0 {
		// Delete the member's message if a filter is violated
		session.ChannelMessageDelete(message.ChannelID, message.ID)

		offenderChannel, _ := session.UserChannelCreate(message.Author.ID)
		guild, _ := session.Guild(message.GuildID)
		author := message.Author.ID

		if !strings.HasPrefix(message.Content, config.CommandPrefix) {
			output := "<@" + author + ">, I detected you said a bad thing and I removed it.\n"
			output += ">>> "
			output += "You said ***" + message.Content + "***\n"

			if violations[0] != message.Content {
				output += "I detected "
				for i := 0; i < len(violations); i++ {
					output += "***" + violations[i] + "***"
					if len(violations) > 1 {
						if i == len(violations)-2 {
							output += " and "
						} else if i < len(violations)-1 {
							output += ", "
						}
					}
				}
				output += " in that.\n"
			}
			output += "Detected in ***" + guild.Name + "*** in the channel <#" + message.ChannelID + ">\n"

			if len(Alerts) > 0 {
				for _, alert := range Alerts {
					alertUser, _ := session.User(alert.ID)
					if !alertUser.Bot {
						alertChannel, err := session.UserChannelCreate(alert.ID)

						if err != nil {
							log.Println(err)
						}

						alertText := "<@" + alert.ID + ">, member <@" + author + "> violated a filter.\n"
						alertText += ">>> "
						alertText += "They said ***" + message.Content + "***\n"

						if violations[0] != message.Content {
							alertText += "I detected "
							for i := 0; i < len(violations); i++ {
								alertText += "***" + violations[i] + "***"
								if len(violations) > 1 {
									if i == len(violations)-2 {
										alertText += " and "
									} else if i < len(violations)-1 {
										alertText += ", "
									}
								}
							}
							alertText += " in that.\n"
						}
						alertText += "Detected in ***" + guild.Name + "*** in the channel <#" + message.ChannelID + ">\n"
						session.ChannelMessageSend(alertChannel.ID, alertText)
					}
				}
				output += "I alerted the appropriate members so they can review this.\n"
			}

			if config.MuteRole != "" {
				// Give the member the mute role if a filter is violated
				session.GuildMemberRoleAdd(message.GuildID, message.Author.ID, config.MuteRole)
				output += "***YOU HAVE BEEN MUTED BECAUSE OF THIS!***\n"
			}

			session.ChannelMessageSend(offenderChannel.ID, output)
		}
	}
}

// LoadFilters and filters alert from json file
func LoadFilters() error {
	log.Println("Loading Filters")

	file, err := ioutil.ReadFile(BinPath + filtersFile)

	// Create new json file if none exists, returns for other errors
	if err != nil {
		if strings.Contains(err.Error(), "no") && strings.Contains(err.Error(), "directory") {
			saveFilters()
			return nil
		}
		return err
	}

	err = json.Unmarshal(file, &Filters)

	// Return if there was an error unmarshaling json
	if err != nil {
		return err
	}

	file, err = ioutil.ReadFile(BinPath + alertsFile)

	// Create new json file if none exists, returns for other errors
	if err != nil {
		if strings.Contains(err.Error(), "no") && strings.Contains(err.Error(), "directory") {
			saveAlerts()
			return nil
		}
		return err
	}

	err = json.Unmarshal(file, &Alerts)

	return nil
}

// SaveFilter to json file
func SaveFilter(text string) error {
	for _, filter := range Filters {
		if filter.Text == text {
			// If the filter already exists return an error
			return errors.New("Filter already exists \"" + text + "\"")
		}
	}

	newFilter := filter{Text: text}
	Filters = append(Filters, newFilter)
	saveFilters()
	return nil
}

// RemoveFilter from json file
func RemoveFilter(text string) error {
	for i, filter := range Filters {
		if filter.Text == text {
			// Remove the filter from Filters
			Filters[i] = Filters[len(Filters)-1]
			Filters = Filters[:len(Filters)-1]

			saveFilters()
			return nil
		}
	}
	// If the filter was not found return an error
	return errors.New("Cannot find filter \"" + text + "\"")
}

// SaveAlert user will be alerted if a mute filter is violated
func SaveAlert(user discordgo.User) error {
	for _, existingUser := range Alerts {
		if user.ID == existingUser.ID {
			// If the filter already exists return an error
			return errors.New("Alert for user \"" + user.ID + "\" already exists")
		}
	}

	newAlert := alert{ID: user.ID}
	Alerts = append(Alerts, newAlert)
	saveAlerts()
	return nil
}

// RemoveAlert from json file
func RemoveAlert(user discordgo.User) error {
	for i, alerts := range Alerts {
		if alerts.ID == user.ID {
			// Remove user from Alerts
			Alerts[i] = Alerts[len(Alerts)-1]
			Alerts = Alerts[:len(Alerts)-1]

			saveAlerts()
			return nil
		}
	}
	// If the filter was not found return an error
	return errors.New("Cannot find alert for user \"" + user.ID + "\"")
}

func saveFilters() error {
	FiltersJSON, err := json.MarshalIndent(Filters, "", " ")

	// Return if there was an error marshaling Filters
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = ioutil.WriteFile(BinPath+filtersFile, FiltersJSON, 0644)

	// Return if there was an error writing to json file
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func saveAlerts() error {
	AlertsJSON, err := json.MarshalIndent(Alerts, "", " ")

	// Return if there was an error marshaling filters
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = ioutil.WriteFile(BinPath+alertsFile, AlertsJSON, 0644)

	// Return if there was an error writing to json file
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
