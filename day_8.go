package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day8Puzzle{})
}

type Day8Puzzle struct{}

func (day *Day8Puzzle) Number() int {
	return 8
}

func (day *Day8Puzzle) Name() string {
	return fmt.Sprintf("day_%02d", day.Number())
}

func (day *Day8Puzzle) Solution() (*Result, error) {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	data, err := i.Strings(2022, day.Number())
	if err != nil {
		return nil, err
	}
	begin := time.Now()

	grid := NewGrid(len(data), len(data[0]))
	for y, row := range data {
		for x, col := range strings.Split(row, "") {
			colVal, _ := strconv.Atoi(col)
			grid.Add(x, y, colVal)
		}
	}
	var puzzleOne = 0
	var puzzleTwo = 0
	for y := range grid.Rows {
		for x := range grid.Rows[y] {
			ok, score := grid.Visibility(x, y)
			if ok {
				puzzleOne++
			}
			if score > puzzleTwo {
				puzzleTwo = score
			}
		}
	}

	return &Result{
		First:    puzzleOne,
		Second:   puzzleTwo,
		Duration: time.Now().Sub(begin),
	}, nil
}

type Grid struct {
	Rows    [][]int
	Columns [][]int
}

func NewGrid(l, w int) *Grid {
	return &Grid{
		Rows:    make([][]int, l),
		Columns: make([][]int, w),
	}
}

func (g *Grid) Add(x, y, val int) {
	if g.Rows[y] == nil {
		g.Rows[y] = make([]int, len(g.Columns))
	}
	if g.Columns[x] == nil {
		g.Columns[x] = make([]int, len(g.Rows))
	}
	g.Rows[y][x] = val
	g.Columns[x][y] = val
}

func (g Grid) Visibility(x, y int) (bool, int) {
	var tree = g.Columns[x][y]

	visLeft, distLeft := visibility(reverse(g.Rows[y][:x]), tree)
	visRight, distRight := visibility(g.Rows[y][x+1:], tree)
	visUp, distUp := visibility(reverse(g.Columns[x][:y]), tree)
	visDown, distDown := visibility(g.Columns[x][y+1:], tree)

	return (visLeft || visRight || visDown || visUp),
		(distLeft * distRight * distDown * distUp)
}

func visibility(trees []int, tree int) (bool, int) {
	distance := 0
	for _, v := range trees {
		distance++
		if v >= tree {
			return false, distance
		}
	}
	return true, distance
}

func reverse(input []int) []int {
	inputLen := len(input)
	output := make([]int, inputLen)

	for i, n := range input {
		j := inputLen - i - 1

		output[j] = n
	}

	return output
}
