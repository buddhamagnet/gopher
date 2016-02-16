package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Printf("machine has %d cores\n", runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	fmt.Println("starting goroutines")
	go printPrime("A")
	go printPrime("B")
	fmt.Println("waiting for goroutines")
	wg.Wait()
	fmt.Println("\nterminating program")
}

func printPrime(label string) {
	defer wg.Done()

next:
	for outer := 2; outer <= 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", label, outer)
	}
	fmt.Println("completed", label)
}
