package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)
	go doWork("A")
	go doWork("B")
	time.Sleep(1 * time.Second)
	fmt.Println("shutdown NOW")
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}

func abort() bool {
	return atomic.LoadInt64(&shutdown) == 1
}

func doWork(label string) {
	defer wg.Done()
	for {
		fmt.Printf("%s is working\n", label)
		time.Sleep(250 * time.Millisecond)
		if abort() {
			fmt.Printf("%s is shutting down\n", label)
			break
		}
	}
}
