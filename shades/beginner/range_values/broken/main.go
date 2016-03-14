package main

// This can happen if you are used to the "for-in" or "foreach" statements in other languages.
// The range clause in Go is different. It generates two values: the first value is the item index while
// the second value is the item data.

import "fmt"

func main() {
	x := []string{"a", "b", "c"}

	for v := range x {
		fmt.Println(v)
	}
}
