package main

import (
	"time"
)

func init() {
	puzzles = append(puzzles, &Day9Puzzle{})
}

type Day9Puzzle struct{}

func (day *Day9Puzzle) Name() string {
	return "day_9"
}

func (day *Day9Puzzle) Solution() (*Result, error) {
	begin := time.Now()
	// i, err := aocutil.NewInputFromFile("session_id")
	// if err != nil {
	// 	return nil, err
	// }

	// data, err := i.Strings(2022, 9)
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
