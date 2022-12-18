package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day13Puzzle{})
}

type Day13Puzzle struct{}

func (day *Day13Puzzle) Number() int {
	return 13
}

func (day *Day13Puzzle) Name() string {
	return fmt.Sprintf("day_%02d", day.Number())
}

func (day *Day13Puzzle) Solution() (*Result, error) {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	data, err := i.Strings(2022, day.Number())
	if err != nil {
		return nil, err
	}
	begin := time.Now()

	packets := make([][]any, 0)
	for _, p := range data[0:2] {
		var packet []any
		if err := json.Unmarshal([]byte(p), &packet); err != nil {
			return nil, err
		}
		packets = append(packets, packet)
	}
	fmt.Println(packets[0][0])
	fmt.Println(packets[1][0])

	return &Result{
		First:    nil,
		Second:   nil,
		Duration: time.Now().Sub(begin),
	}, nil
}
