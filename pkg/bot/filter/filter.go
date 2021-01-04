package filter

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strings"
)

var (
	// BinPath is the location of the executable binary
	BinPath string

	// Filters that members will be muted for messaging
	Filters []Filter

	filtersFile       = "/json/filters.json"
	filtersNotifyFile = "/json/filtersnotify.json"
)

// Filter phrase or word members will be muted for messaging
type Filter struct {
	Text string `json:"text"`
}

// LoadFilters and filters notify from filters.json and filtersnotify.json
func LoadFilters() error {
	log.Println("Loading filters")

	file, err := ioutil.ReadFile(BinPath + filtersFile)

	// Create new reactionroles.json file if none exists, returns for other errors
	if err != nil {
		if strings.Contains(err.Error(), "no") && strings.Contains(err.Error(), "directory") {
			saveFilters()
			return nil
		}
		return err
	}

	err = json.Unmarshal(file, &Filters)

	// Return if there was an error unmarshaling reactionrroles.json
	if err != nil {
		return err
	}

	return nil
}

// SaveFilter to filter.json
func SaveFilter(text string) error {
	for _, filter := range Filters {
		if filter.Text == text {
			// If the filter already exists return an error.
			return errors.New("Filter already exists \"" + text + "\"")
		}
	}

	newFilter := Filter{Text: text}
	Filters = append(Filters, newFilter)
	saveFilters()
	return nil
}

// RemoveFilter from filter.json
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
	// If the filter was not found return an error.
	return errors.New("Cannot find filter \"" + text + "\"")
}

func saveFilters() error {
	filtersJSON, err := json.MarshalIndent(Filters, "", " ")

	// Return if there was an error marshaling Filters
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = ioutil.WriteFile(BinPath+filtersFile, filtersJSON, 0644)

	// Return if there was an error writing to filters.json
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
