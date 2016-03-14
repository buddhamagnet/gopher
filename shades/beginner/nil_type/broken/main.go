package main

// The nil identifier can be used as the zero value for interfaces, functions, pointers, maps, slices
// and channels. If you don't specify the variable type the compiler will fail to compile
// your code because it can't guess the type.
func main() {
	// computer says use of untyped nil.
	var x = nil
}
