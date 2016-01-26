package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Register(name string, matcher Matcher) {
	if _, exists := matchers[name]; exists {
		log.Fatalln(name, "matcher already registered")
	}
	log.Println(name, "matcher successfully registered")
	matchers[name] = matcher
}

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
