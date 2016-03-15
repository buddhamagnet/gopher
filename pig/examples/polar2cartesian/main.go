package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type polar struct {
	radius float64
	m      float64
}

type cartesian struct {
	x float64
	y float64
}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			polarCoord := <-questions
			m := polarCoord.m * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(m)
			y := polarCoord.radius * math.Sin(m)
			answers <- cartesian{x, y}
		}
	}()
	return answers
}

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter a radius and an angle or ctrl+d to quit")
	for {
		fmt.Printf("radius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var radius, m float64
		if _, err := fmt.Sscanf(line, "%f %f", &radius, &m); err != nil {
			fmt.Fprintln(os.Stderr, "invalid input")
			continue
		}
		questions <- polar{radius, m}
		coord := <-answers
		fmt.Printf("Polar radius: %.02f m: %.02f Cartesian x: %.02f y: %.02f\n", radius, m, coord.x, coord.y)
		fmt.Println()
	}
}
