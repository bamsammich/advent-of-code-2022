package main

import (
	"time"
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
	// i, err := aocutil.NewInputFromFile("session_id")
	// if err != nil {
	// 	return nil, err
	// }

	// data, err := i.Strings(2022, 6)
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println(data)

	return &Result{
		First:    nil,
		Second:   nil,
		Duration: time.Now().Sub(begin),
	}, nil
}
