package main

import (
	"time"
)

func init() {
	puzzles = append(puzzles, &Day8Puzzle{})
}

type Day8Puzzle struct{}

func (day *Day8Puzzle) Name() string {
	return "day_8"
}

func (day *Day8Puzzle) Solution() (*Result, error) {
	begin := time.Now()
	// i, err := aocutil.NewInputFromFile("session_id")
	// if err != nil {
	// 	return nil, err
	// }

	// data, err := i.Strings(2022, 8)
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
