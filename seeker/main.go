package main

import (
	"flag"
	"log"
	"os"

	_ "github.com/buddhamagnet/gia/seeker/matchers"
	"github.com/buddhamagnet/gia/seeker/search"
)

var term string

func init() {
	flag.StringVar(&term, "term", "golang", "search term")
	log.SetOutput(os.Stdout)
}

func main() {
	flag.Parse()
	search.Run(term)
}
