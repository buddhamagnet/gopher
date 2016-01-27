package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

// Register adds a matcher value to the matcher map.
func Register(name string, matcher Matcher) {
	if _, exists := matchers[name]; exists {
		log.Fatalln(name, "matcher already registered")
	}
	log.Println(name, "matcher successfully registered")
	matchers[name] = matcher
}

// Run takes a search term, pulls in the feeds and
// fires up and synchronises the goroutines and channels
// that process and return the search results.
func Run(term string) {
	var wg sync.WaitGroup
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	results := make(chan *Result)
	wg.Add(len(feeds))
	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}
		go func(matcher Matcher, feed *Feed) {
			defer wg.Done()
			Match(matcher, feed, term, results)
		}(matcher, feed)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	Display(results)
}
