package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	fmt.Println("start goroutines")

	go func() {
		defer wg.Done()

		for char := 'A'; char < 'A'+26; char++ {
			fmt.Printf("%c", char)
		}
	}()

	go func() {
		defer wg.Done()
		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c", char)
		}
	}()

	fmt.Println("waiting to finish")
	wg.Wait()
	fmt.Println("\nterminating program")
}
