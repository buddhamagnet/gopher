package main

import (
	"log"
	"net/http"

	"github.com/buddhamagnet/gopher/pig/examples/stats/web"
)

func main() {
	web.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", nil))
}
