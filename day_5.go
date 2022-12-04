package main

import (
	"fmt"
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day5Puzzle{})
}

type Day5Puzzle struct{}

func (day *Day5Puzzle) Name() string {
	return "day_5"
}

func (day *Day5Puzzle) Solution() (*Result, error) {
	begin := time.Now()
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	data, err := i.Strings(2022, 4)
	if err != nil {
		return nil, err
	}

	fmt.Println(data)
	return &Result{
		First:    nil,
		Second:   nil,
		Duration: time.Now().Sub(begin),
	}, nil
}
