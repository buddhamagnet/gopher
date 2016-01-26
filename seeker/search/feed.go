package search

import (
	"encoding/json"
	"os"
)

type Feed struct {
	Name string `json:"site"`
	Type string `json:"type"`
	URI  string `json:"link"`
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open("data/data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
