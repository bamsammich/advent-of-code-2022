package main

import (
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day6Puzzle{})
}

type Day6Puzzle struct{}

func (day *Day6Puzzle) Name() string {
	return "day_6"
}

func (day *Day6Puzzle) Solution() (*Result, error) {
	begin := time.Now()
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}
	data, err := i.Strings(2022, 6)
	if err != nil {
		return nil, err
	}
	stream := data[0]

	var puzzleOne int
	for i := 0; i+3 < len(stream); i++ {
		packet := stream[i : i+4]
		if unique([]rune(packet)) {
			puzzleOne = i + 4
			break
		}
	}

	var puzzleTwo int
	for i := 0; i+13 < len(stream); i++ {
		packet := stream[i : i+14]
		if unique([]rune(packet)) {
			puzzleTwo = i + 14
			break
		}
	}

	return &Result{
		First:    puzzleOne,
		Second:   puzzleTwo,
		Duration: time.Now().Sub(begin),
	}, nil
}

func unique[T comparable](sl []T) bool {
	uniq := make(map[T]bool)
	for _, e := range sl {
		if ok := uniq[e]; ok {
			return false
		}
		uniq[e] = true
	}
	return true
}
