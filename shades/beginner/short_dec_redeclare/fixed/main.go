package main

import "fmt"

func main() {
	x := "hello"
	if true {
		// fixed - run go tool vet -shadow=true to catch.
		x := "shadow!"
		fmt.Println(x)
	}
	newX(x)
}

func newX(message string) {
	x := "fixed!"
	fmt.Println(x)
}
