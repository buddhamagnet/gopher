package main

import "fmt"

func main() {
	// fixed - interfaces, functions, pointers, maps and slices can
	// all be assigned nil. This is for explicit demonstration only,
	// simply declaring them will initialize them to their default
	// value, which is nil.
	var i interface{} = nil
	var f func() = nil
	var p *int = nil
	var m map[string]int = nil
	var s []string = nil
	var c chan int = nil

	fmt.Println(i)
	// Calling f will cause a runtime panic - invalid memory address or nil pointer dereference.
	fmt.Println(f)
	fmt.Println(p)
	// Attempting to add an entry to the map will cause a runtime panic - assignment to entry in nil map.
	// Attempting to retrieve an entry will return the default value for the map element type.
	fmt.Println(m)
	fmt.Println(m["test"])
	// Attempting to read an element from the slice will cause a runtime panic - index out of range.
	fmt.Println(s)
	fmt.Println(c)
}
