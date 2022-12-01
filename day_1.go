package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day1Puzzle{})
}

type Day1Puzzle struct{}

func (day *Day1Puzzle) Name() string {
	return "Day 1"
}

func (day *Day1Puzzle) Solution() (any, error) {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	data, err := i.Strings(2022, 1)
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
	return struct {
		one, two int
	}{
		one: elfCals[len(elfCals)-1],
		two: sumInts(elfCals[len(elfCals)-3 : len(elfCals)]),
	}, nil
}

func sumInts(ints []int) int {
	var sum = 0
	for _, i := range ints {
		sum += i
	}
	return sum
}
