package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	court := make(chan int)
	wg.Add(2)
	go player("dave", court)
	go player("rachel", court)
	court <- 1
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		time.Sleep(1 * time.Second)
		ball, ok := <-court
		if !ok {
			fmt.Printf("player %s won\n", name)
			return
		}
		m := rand.Intn(100)
		if m%13 == 0 {
			fmt.Printf("player %s missed\n", name)
			close(court)
			return
		}
		fmt.Printf("player %s hit the ball %d times\n", name, ball)
		ball++
		court <- ball
	}
}
