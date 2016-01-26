package search

import (
	"fmt"
	"log"
)

type Result struct {
	Field   string
	Content string
}

func (r Result) String() string {
	return fmt.Sprintf("%s:\n%s\n\n", r.Field, r.Content)
}

type Matcher interface {
	Search(feed *Feed, term string) ([]*Result, error)
}

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

func Display(results chan *Result) {
	for result := range results {
		fmt.Println(result)
	}

}
