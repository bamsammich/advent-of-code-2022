package main

import (
	"fmt"
	"time"
)

func init() {
	puzzles = append(puzzles, &Day12Puzzle{})
}

type Day12Puzzle struct{}

func (day *Day12Puzzle) Number() int {
	return 12
}

func (day *Day12Puzzle) Name() string {
	return fmt.Sprintf("day_%02d", day.Number())
}

func (day *Day12Puzzle) Solution() (*Result, error) {
	begin := time.Now()
	// i, err := aocutil.NewInputFromFile("session_id")
	// if err != nil {
	// 	return nil, err
	// }

	// data, err := i.Strings(2022, day.Number())
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
