package main

import (
	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day2Puzzle{})
}

type Day2Puzzle struct{}

func (day *Day2Puzzle) Name() string {
	return "Day 2"
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

func (day *Day2Puzzle) Solution() (any, error) {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	data, err := i.Strings(2022, 2)
	if err != nil {
		return nil, err
	}
	scorePartTwo, scorePartOne := 0, 0
	for _, round := range data {
		scorePartOne += scoresPartOne[round]
		scorePartTwo += scoresPartTwo[round]
	}

	return struct {
		first  int
		second int
	}{
		first:  scorePartOne,
		second: scorePartTwo,
	}, nil
}
