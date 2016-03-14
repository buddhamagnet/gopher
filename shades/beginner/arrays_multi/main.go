package main

import "fmt"

/* Creating a dynamic multi-dimensional array using slices
of "independent" slices is a two step process.
First, you have to create the outer slice. Then, you have
to allocate each inner slice. The inner slices are
independent of each other. You can grow and shrink them
without affecting other inner slices. */
func main() {
	x := 2
	y := 4

	table := make([][]int, x)
	for i := range table {
		table[i] = make([]int, y)
	}
	fmt.Println(len(table))
	for _, j := range table {
		fmt.Println(len(j))
	}
}
