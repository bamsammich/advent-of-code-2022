package main

import (
	"fmt"
	"time"
)

func init() {
	puzzles = append(puzzles, &Day15Puzzle{})
}

type Day15Puzzle struct{}

func (day *Day15Puzzle) Number() int {
	return 15
}

func (day *Day15Puzzle) Name() string {
	return fmt.Sprintf("day_%02d", day.Number())
}

func (day *Day15Puzzle) Solution() (*Result, error) {
	// i, err := aocutil.NewInputFromFile("session_id")
	// if err != nil {
	// 	return nil, err
	// }

	// data, err := i.Strings(2022, day.Number())
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(data)
	begin := time.Now()

	return &Result{
		First:    nil,
		Second:   nil,
		Duration: time.Now().Sub(begin),
	}, nil
}
