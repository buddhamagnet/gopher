package main

import "fmt"

// Also, unused package level variables are acceptable.
var global string

// Also, unused function arguments are OK.
func dump(message string, code int) {
	fmt.Println(message)
}

func main() {
	x := "hello"
	// Fixed, use me and I'll smile.
	// So, what we are really picky about are variables
	// declared inside functions that are not used.
	dump(x, 100)
}
