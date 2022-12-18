package main

import (
	"fmt"
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day3Puzzle{})
}

type Day3Puzzle struct{}

func (day *Day3Puzzle) Number() int {
	return 3
}

func (day *Day3Puzzle) Name() string {
	return fmt.Sprintf("day_%02d", day.Number())
}

var letterValues = map[rune]int{
	'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10, 'k': 11, 'l': 12, 'm': 13,
	'n': 14, 'o': 15, 'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20, 'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26,
	'A': 27, 'B': 28, 'C': 29, 'D': 30, 'E': 31, 'F': 32, 'G': 33, 'H': 34, 'I': 35, 'J': 36, 'K': 37, 'L': 38, 'M': 39,
	'N': 40, 'O': 41, 'P': 42, 'Q': 43, 'R': 44, 'S': 45, 'T': 46, 'U': 47, 'V': 48, 'W': 49, 'X': 50, 'Y': 51, 'Z': 52,
}

func (day *Day3Puzzle) Solution() (*Result, error) {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	data, err := i.Strings(2022, day.Number())
	if err != nil {
		return nil, err
	}
	begin := time.Now()

	var firstPuzzle = 0
	for _, ruck := range data {
		compartmentOne := ruck[:len(ruck)/2]
		compartmentTwo := ruck[len(ruck)/2:]
		for _, item := range []rune(compartmentOne) {
			if contains([]rune(compartmentTwo), item) {
				firstPuzzle += letterValues[item]
				break
			}
		}
	}
	var secondPuzzle = 0
	for i := 2; i <= len(data); i += 3 {
		for _, item := range []rune(data[i]) {
			if contains([]rune(data[i-1]), item) && contains([]rune(data[i-2]), item) {
				secondPuzzle += letterValues[item]
				break
			}
		}
	}
	return &Result{
		First:    firstPuzzle,
		Second:   secondPuzzle,
		Duration: time.Now().Sub(begin),
	}, nil
}

func contains[T comparable](sl []T, item T) bool {
	for _, e := range sl {
		if e == item {
			return true
		}
	}
	return false
}
