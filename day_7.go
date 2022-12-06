package main

import (
	"time"
)

func init() {
	puzzles = append(puzzles, &Day7Puzzle{})
}

type Day7Puzzle struct{}

func (day *Day7Puzzle) Name() string {
	return "day_7"
}

func (day *Day7Puzzle) Solution() (*Result, error) {
	begin := time.Now()
	// i, err := aocutil.NewInputFromFile("session_id")
	// if err != nil {
	// 	return nil, err
	// }

	// data, err := i.Strings(2022, 7)
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
