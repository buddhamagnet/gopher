package search

func RetrieveFeeds() (feeds []*Feed, err error) {
	return []*Feed{}, nil
}

func Match(m Matcher, feed *Feed, term string, results chan *Result) {
}

func Display(results chan *Result) {
}

type Result struct {
}

type Feed struct {
	Type string
}
