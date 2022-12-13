package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day1Puzzle{})
}

var dayNumber = 1

type Day1Puzzle struct{}

func (day *Day1Puzzle) Number() int {
	return 1
}

func (day *Day1Puzzle) Name() string {
	return fmt.Sprintf("day_%02d", day.Number())
}

func (day *Day1Puzzle) Solution() (*Result, error) {
	begin := time.Now()
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	data, err := i.Strings(2022, dayNumber)
	if err != nil {
		return nil, err
	}
	var (
		elfCals []int
		calSum  = 0
	)
	for _, calStr := range data {
		if calStr == "" {
			elfCals = append(elfCals, calSum)
			calSum = 0
			continue
		}
		cal, err := strconv.Atoi(calStr)
		if err != nil {
			return nil, fmt.Errorf("failed to convert %s to number: %v", calStr, err)
		}
		calSum += cal
	}
	sort.Ints(elfCals)
	return &Result{
		First:    elfCals[len(elfCals)-1],
		Second:   sumInts(elfCals[len(elfCals)-3:]),
		Duration: time.Now().Sub(begin),
	}, nil
}
