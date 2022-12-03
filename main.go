package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Puzzle interface {
	Name() string
	Solution() (*Result, error)
}

type Result struct {
	First    any
	Second   any
	Duration time.Duration
}

func (r *Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		First    any    `json:"first"`
		Second   any    `json:"second"`
		Duration string `json:"duration"`
	}{
		First:    r.First,
		Second:   r.Second,
		Duration: r.Duration.String(),
	})
}

var puzzles []Puzzle

func main() {
	var totalDuration int64
	results := make(map[string]*Result)
	for _, p := range puzzles {
		solution, err := p.Solution()
		if err != nil {
			log.Fatalf("failed to solve puzzle %s: %v", p.Name(), err)
			continue
		}
		results[p.Name()] = solution
		totalDuration += solution.Duration.Nanoseconds()
	}
	fmt.Println(PrettyJSON(&results))
	fmt.Printf("Total duration: %s\n", time.Duration(totalDuration))
}
