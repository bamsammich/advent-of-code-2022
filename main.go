package main

import (
	"fmt"
	"log"
)

type Puzzle interface {
	Name() string
	Solution() (any, error)
}

var puzzles []Puzzle

func main() {
	for _, p := range puzzles {
		solution, err := p.Solution()
		if err != nil {
			log.Fatalf("failed to solve puzzle %s: %v", p.Name(), err)
			continue
		}
		fmt.Printf("Puzzle %s: %+v\n", p.Name(), solution)
	}
}
