package main

import (
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day4Puzzle{})
}

type Day4Puzzle struct{}

func (day *Day4Puzzle) Name() string {
	return "day_4"
}

func (day *Day4Puzzle) Solution() (*Result, error) {
	begin := time.Now()
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	data, err := i.Strings(2022, 4)
	if err != nil {
		return nil, err
	}

	return &Result{
		First:    nil,
		Second:   nil,
		Duration: time.Now().Sub(begin),
	}, nil
}
