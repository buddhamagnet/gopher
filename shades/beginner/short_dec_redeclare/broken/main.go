package main

func main() {
	x := "hello"
	// broken - can't redeclare with short variable declaration.
	// computer says no new variables on left side of :=.
	x := "world"

}
