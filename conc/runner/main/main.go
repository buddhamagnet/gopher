package main

import (
	"log"
	"os"
	"time"

	"github.com/buddhamagnet/gopher/conc/runner"
)

const timeout = 3 * time.Second

func main() {
	log.Println("starting work")
	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("terminating due to timeout\n")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("terminating due to interrupt")
			os.Exit(2)
		}
	}
	log.Println("process ended")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("processor task %d\n", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
