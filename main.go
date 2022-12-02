package main

import (
	"fmt"
	"log"
)

type Puzzle interface {
	Name() string
	Solution() (*Result, error)
}

type Result struct {
	First  any `json:"first" yaml:"first"`
	Second any `json:"second" yaml:"second"`
}

var puzzles []Puzzle

func main() {
	results := make(map[string]*Result)
	for _, p := range puzzles {
		solution, err := p.Solution()
		if err != nil {
			log.Fatalf("failed to solve puzzle %s: %v", p.Name(), err)
			continue
		}
		results[p.Name()] = solution
	}
	fmt.Println(PrettyJSON(&results))
}
