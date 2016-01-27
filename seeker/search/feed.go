package search

import (
	"encoding/json"
	"fmt"
	"os"
)

// Feed represents an RSS feed in the program.
type Feed struct {
	Name string `json:"site"`
	Type string `json:"type"`
	URI  string `json:"link"`
}

// String is the Feed stringer implementation.
func (f Feed) String() string {
	return fmt.Sprintf("feed type: %s site: %s URI: %s\n", f.Type, f.Name, f.URI)
}

// RetrieveFeeds returns a slice of Feed values from
// a data file.
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
