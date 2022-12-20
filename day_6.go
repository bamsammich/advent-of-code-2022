package main

import (
	"fmt"
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day6Puzzle{})
}

type Day6Puzzle struct{}

func (day *Day6Puzzle) Number() int {
	return 6
}

func (day *Day6Puzzle) Name() string {
	return fmt.Sprintf("day_%02d", day.Number())
}

func (day *Day6Puzzle) Solution() (*Result, error) {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}
	data, err := i.Strings(2022, day.Number())
	if err != nil {
		return nil, err
	}
	begin := time.Now()

	stream := data[0]

	var puzzleOne int
	for i := 0; i+3 < len(stream); i++ {
		packet := stream[i : i+4]
		if isUnique([]rune(packet)) {
			puzzleOne = i + 4
			break
		}
	}

	var puzzleTwo int
	for i := 0; i+13 < len(stream); i++ {
		packet := stream[i : i+14]
		if isUnique([]rune(packet)) {
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
