package search

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

type defaultMatcher struct{}

// Search is the defaultMatcher Matcher implementation.
func (m defaultMatcher) Search(feed *Feed, term string) ([]*Result, error) {
	return nil, nil
}
