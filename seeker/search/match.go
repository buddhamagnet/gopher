package search

import (
	"fmt"
	"log"
)

// Result represents a search result.
type Result struct {
	Field   string
	Content string
}

// NewResult is the Result factory function.
func NewResult(field, content string) *Result {
	return &Result{Field: field, Content: content}
}

// String is the Result stringer implementation.
func (r Result) String() string {
	return fmt.Sprintf("%s:\n%s\n\n", r.Field, r.Content)
}

// Matcher is the matching interface.
type Matcher interface {
	Search(feed *Feed, term string) ([]*Result, error)
}

// Match takes search results from a matcher and puts them
// into a results channel.
func Match(matcher Matcher, feed *Feed, term string, results chan *Result) {
	searchResults, err := matcher.Search(feed, term)
	if err != nil {
		log.Println(err)
		return
	}
	for _, result := range searchResults {
		results <- result
	}
}

// Display displays the search results.
func Display(results chan *Result) {
	for result := range results {
		fmt.Println(result)
	}

}
