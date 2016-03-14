package main

import "fmt"

// You'd never do this because slices.

func main() {
	x := [3]int{1, 2, 3}

	func(arr *[3]int) {
		arr[0] = 7
		fmt.Println(*arr)
	}(&x)

	fmt.Println(x)
}
