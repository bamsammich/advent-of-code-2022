package main

import (
	"fmt"
	"time"

	"github.com/echojc/aocutil"
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

var lowercase = []rune{
	'S', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
	'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'E',
}

func (day *Day12Puzzle) Solution() (*Result, error) {
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
		start         = Position{0, 20}
		end           = Position{137, 20}
		graph         = NewNodeGraph()
		lowElevations []Position
	)
	for y := range data {
		for x, ch := range data[y] {
			var (
				pos       = Position{x, y}
				edges     = graph.GetEdges(pos)
				neighbors = pos.CardinalNeighbors()
			)
			if ch == 'a' {
				lowElevations = append(lowElevations, pos)
			}
			for _, n := range neighbors {
				if n.X < 0 || n.Y < 0 || n.X > len(data[0])-1 || n.Y > len(data)-1 {
					continue // out of bounds
				}
				var (
					nLetter = []rune(data[n.Y])[n.X]
					nVal    = index(lowercase, nLetter)
					posVal  = index(lowercase, ch)
					weight  = (nVal - posVal) + 1
				)
				if weight > 2 {
					continue
				}
				if !contains(edges, Edge{Node: n, Weight: weight}) {
					if nLetter == 'c' && ch == 'a' {
						fmt.Printf("%c(%d) - %c(%d) = %d\n", nLetter, nLetter, ch, posVal, weight)
					}
					graph.AddEdge(pos, n, weight)
				}
			}
		}
	}

	_, puzzleOne := graph.GetPath(start, end)

	var puzzleTwo []Position
	for _, e := range lowElevations {
		_, path := graph.GetPath(e, end)
		if len(puzzleTwo) == 0 {
			puzzleTwo = path
			continue
		}
		if len(path) > 0 && len(path) < len(puzzleTwo) {
			puzzleTwo = path
		}
	}
	return &Result{
		First:    len(puzzleOne) - 1,
		Second:   len(puzzleTwo) - 1,
		Duration: time.Now().Sub(begin),
	}, nil
}

func printPath(path []Position, data []string) {
	colorRed := "\033[31m"
	colorReset := "\033[0m"
	for y := range data {
		for x, ch := range data[y] {
			if contains(path, Position{x, y}) {
				fmt.Printf("%s%c%s", colorRed, ch, colorReset)
			} else {
				fmt.Print(string(ch))
			}
		}
		fmt.Print("\n")
	}
}
