package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	after := time.After(10 * time.Second)
	stream := make(chan string)
	wg.Add(100)
	for i := 0; i <= 99; i++ {
		go spawn(i, stream, after)
	}
	for j := 0; j <= 1000; j++ {
		stream <- strconv.Itoa(j) + "_string"
	}
	wg.Wait()
	fmt.Println("all go funcs stopped")
}

func spawn(id int, stream chan string, after <-chan time.Time) {
	defer wg.Done()
	for {
		select {
		case <-after:
			fmt.Printf("shutdown signal received by go func %d\n", id)
			close(stream)
			return
		case name, ok := <-stream:
			if !ok {
				fmt.Printf("go func %d reports stream drained\n", id)
				return
			}
			fmt.Printf("go func %d says %s\n", id, name)
			time.Sleep(1 * time.Second)
		}
	}
}
