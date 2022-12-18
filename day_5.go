package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day5Puzzle{})
}

type Day5Puzzle struct{}

func (day *Day5Puzzle) Number() int {
	return 5
}

func (day *Day5Puzzle) Name() string {
	return fmt.Sprintf("day_%02d", day.Number())
}

type day5Move struct {
	Count, Source, Destination int
}

func (day *Day5Puzzle) Solution() (*Result, error) {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	data, err := i.Strings(2022, day.Number())
	if err != nil {
		return nil, err
	}
	begin := time.Now()

	stacks := make([][]string, 9)
	for i := 7; i >= 0; i-- {
		for r := 0; r < len(data[i]); r += 4 {
			crate := string(data[i][r : r+3])
			if crate == "   " {
				continue
			}
			stackNum := r / 4
			if stacks[stackNum] == nil {
				stacks[stackNum] = []string{}
			}
			stacks[stackNum] = append(stacks[stackNum], crate)
		}
	}

	var moves []day5Move
	for _, move := range data[10:] {
		match := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`).FindAllStringSubmatch(move, 1)[0]
		count, _ := strconv.Atoi(match[1])
		src, _ := strconv.Atoi(match[2])
		dest, _ := strconv.Atoi(match[3])
		moves = append(moves, day5Move{count, src - 1, dest - 1})
	}

	return &Result{
		First:    day5PuzzleOne(copyStacks(stacks), moves),
		Second:   day5PuzzleTwo(copyStacks(stacks), moves),
		Duration: time.Now().Sub(begin),
	}, nil
}

func day5PuzzleOne(stacks [][]string, moves []day5Move) string {
	for _, move := range moves {
		for i := 0; i < move.Count; i++ {
			var (
				src  = move.Source
				dest = move.Destination
			)
			stacks[dest] = append(stacks[dest], stacks[src][len(stacks[src])-1])
			stacks[src] = stacks[src][:len(stacks[src])-1]
		}
	}
	var output string
	for i := range stacks {
		output += strings.NewReplacer("[", "", "]", "").Replace(stacks[i][len(stacks[i])-1])
	}
	return output
}

func day5PuzzleTwo(stacks [][]string, moves []day5Move) string {
	for _, move := range moves {
		var (
			src  = move.Source
			dest = move.Destination
		)
		stacks[dest] = append(stacks[dest], stacks[src][len(stacks[src])-move.Count:]...)
		stacks[src] = stacks[src][:len(stacks[src])-move.Count]
	}

	var output string
	for i := range stacks {
		output += strings.NewReplacer("[", "", "]", "").Replace(stacks[i][len(stacks[i])-1])
	}
	return output
}

func copyStacks(stacks [][]string) [][]string {
	var out = make([][]string, len(stacks))
	for i := range stacks {
		out[i] = []string{}
		for e := range stacks[i] {
			out[i] = append(out[i], stacks[i][e])
		}
	}
	return out
}
