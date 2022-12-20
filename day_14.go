package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day14Puzzle{})
}

type Day14Puzzle struct{}

func (day *Day14Puzzle) Number() int {
	return 14
}

func (day *Day14Puzzle) Name() string {
	return fmt.Sprintf("day_%02d", day.Number())
}

func (day *Day14Puzzle) Solution() (*Result, error) {
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
		maxX, maxY   int
		rockPoints   []Position
		filledPoints = make(map[Position]bool)
	)
	for _, rockPath := range data {
		points := []Position{}
		rockCoord := strings.Split(rockPath, " -> ")
		for _, p := range rockCoord {
			xy := strings.Split(p, ",")
			x, _ := strconv.Atoi(xy[0])
			if x > maxX {
				maxX = x
			}
			y, _ := strconv.Atoi(xy[1])
			if y > maxY {
				maxY = y
			}
			points = append(points, Position{x, y})
		}

		for i := 1; i < len(points); i++ {
			line := line(points[i-1], points[i])
			rockPoints = append(rockPoints, line...)
		}
	}
	rockPoints = append(rockPoints, line(Position{0, maxY + 2}, Position{maxX + 400, maxY + 2})...)
	for _, p := range rockPoints {
		filledPoints[p] = true
	}

	var (
		sandPoints []Position
		puzzleOne  int
		grain      = Position{500, 0}
	)

FALL:
	for true {
		nextPositions := []Position{
			{grain.X, grain.Y + 1},     // down
			{grain.X - 1, grain.Y + 1}, // down-left
			{grain.X + 1, grain.Y + 1}, // down-right
		}
		for _, next := range nextPositions {
			if !filledPoints[next] {
				grain = next
				continue FALL
			}
		}
		sandPoints = append(sandPoints, grain)
		filledPoints[grain] = true
		if filledPoints[Position{500, 0}] {
			break
		}
		if puzzleOne == 0 && grain.Y > maxY {
			puzzleOne = len(sandPoints) - 1
		}
		grain = Position{500, 0}
	}
	return &Result{
		First:    puzzleOne,
		Second:   len(sandPoints),
		Duration: time.Now().Sub(begin),
	}, nil
}

func line(a, b Position) []Position {
	var (
		sx  = -1
		sy  = -1
		out []Position
	)
	if a.X < b.X {
		sx = 1
	}
	if a.Y < b.Y {
		sy = 1
	}

	switch {
	// point
	case a.X == b.X && a.Y == b.Y:
		out = []Position{a}
	// vertical line
	case a.X == b.X:
		for y := a.Y; y != b.Y+sy; y += sy {
			out = append(out, Position{a.X, y})
		}
	// horizontal line
	case a.Y == b.Y:
		for x := a.X; x != b.X+sx; x += sx {
			out = append(out, Position{x, a.Y})
		}
	}
	return out
}
