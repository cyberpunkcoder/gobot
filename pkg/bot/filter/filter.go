package filter

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	// BinPath is the location of the executable binary
	BinPath string

	// Filters that members will be muted for messaging
	Filters []Filter

	filtersFile = "/json/filters.json"
)

// Filter phrase or word members will be muted for messaging
type Filter struct {
	Value string `json:"value"`
}

// SaveFilter to filter.json
func SaveFilter(value string) {
	newFilter := Filter{Value: value}
	Filters = append(Filters, newFilter)
	saveFilters()
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
