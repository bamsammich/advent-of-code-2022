package main

import (
	"fmt"
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day2Puzzle{})
}

type Day2Puzzle struct{}

func (day *Day2Puzzle) Number() int {
	return 2
}

func (day *Day2Puzzle) Name() string {
	return fmt.Sprintf("day_%02d", day.Number())
}

var scoresPartOne = map[string]int{
	"A X": 4, "A Y": 8, "A Z": 3,
	"B X": 1, "B Y": 5, "B Z": 9,
	"C X": 7, "C Y": 2, "C Z": 6,
}

var scoresPartTwo = map[string]int{
	"A X": 3, "A Y": 4, "A Z": 8,
	"B X": 1, "B Y": 5, "B Z": 9,
	"C X": 2, "C Y": 6, "C Z": 7,
}

func (day *Day2Puzzle) Solution() (*Result, error) {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	data, err := i.Strings(2022, day.Number())
	if err != nil {
		return nil, err
	}
	begin := time.Now()

	scorePartTwo, scorePartOne := 0, 0
	for _, round := range data {
		scorePartOne += scoresPartOne[round]
		scorePartTwo += scoresPartTwo[round]
	}

	return &Result{
		First:    scorePartOne,
		Second:   scorePartTwo,
		Duration: time.Now().Sub(begin),
	}, nil
}
