package main

import "fmt"

var s []string
var m map[string]string

func main() {
	s = append(s, "dave", "is", "amazing")
	fmt.Println(s)
	m["dave"] = "amazing"
}
