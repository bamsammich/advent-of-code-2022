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
	begin := time.Now()
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	data, err := i.Strings(2022, day.Number())
	if err != nil {
		return nil, err
	}

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

type Position struct {
	X, Y int
}

func (pos *Position) Move(x, y int) {
	pos.X += x
	pos.Y += y
}

func (pos *Position) Touches(other *Position) bool {
	xDiff, yDiff := pos.Diff(other)
	return xDiff >= -1 && xDiff <= 1 && yDiff >= -1 && yDiff <= 1
}

func (pos *Position) Diff(from *Position) (int, int) {
	return from.X - pos.X, from.Y - pos.Y
}

func (pos *Position) Direction(to *Position) (dir string) {
	xDiff, yDiff := pos.Diff(to)
	if yDiff > 0 {
		dir = "N"
	} else if yDiff < 0 {
		dir = "S"
	}
	if xDiff > 0 {
		dir = fmt.Sprintf("%sE", dir)
	} else if xDiff < 0 {
		dir = fmt.Sprintf("%sW", dir)
	}
	return
}

func (pos *Position) String() string {
	return fmt.Sprintf("{%d, %d}", pos.X, pos.Y)
}

func (pos *Position) Toward(last *Position) {
	if pos.Touches(last) {
		return
	}
	x, y := 0, 0
	dir := []rune(pos.Direction(last))
	if contains(dir, 'N') {
		y++
	} else if contains(dir, 'S') {
		y--
	}
	if contains(dir, 'E') {
		x++
	} else if contains(dir, 'W') {
		x--
	}
	pos.Move(x, y)
}
