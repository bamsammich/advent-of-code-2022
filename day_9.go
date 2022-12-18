package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day9Puzzle{})
}

type Day9Puzzle struct{}

func (day *Day9Puzzle) Number() int {
	return 9
}

func (day *Day9Puzzle) Name() string {
	return fmt.Sprintf("day_%02d", day.Number())
}

func (day *Day9Puzzle) Solution() (*Result, error) {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	data, err := i.Strings(2022, day.Number())
	if err != nil {
		return nil, err
	}
	begin := time.Now()

	var (
		rope      = make([]Position, 10)
		puzzleOne = make(map[Position]bool)
		puzzleTwo = make(map[Position]bool)
	)
	for _, move := range data {
		dir := strings.Split(move, " ")[0]
		dist, _ := strconv.Atoi(strings.Split(move, " ")[1])
		for i := 0; i < dist; i++ {
			switch dir {
			case "U":
				rope[0].Move(0, 1)
			case "D":
				rope[0].Move(0, -1)
			case "R":
				rope[0].Move(1, 0)
			case "L":
				rope[0].Move(-1, 0)
			}
			for i := range rope {
				if i == 0 {
					continue
				}
				rope[i].Toward(&rope[i-1])
			}
			puzzleOne[rope[1]] = true
			puzzleTwo[rope[9]] = true
		}
	}
	return &Result{
		First:    len(puzzleOne),
		Second:   len(puzzleTwo),
		Duration: time.Now().Sub(begin),
	}, nil
}
