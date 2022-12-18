package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day13Puzzle{})
}

type Day13Puzzle struct{}

func (day *Day13Puzzle) Number() int {
	return 13
}

func (day *Day13Puzzle) Name() string {
	return fmt.Sprintf("day_%02d", day.Number())
}

func (day *Day13Puzzle) Solution() (*Result, error) {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	b, err := i.Bytes(2022, day.Number())
	if err != nil {
		return nil, err
	}
	begin := time.Now()

	var (
		packets   = []any{}
		puzzleOne = 0
	)
	for i, pair := range strings.Split(string(b), "\n\n") {
		pairPackets := strings.Split(pair, "\n")
		var left, right any
		json.Unmarshal([]byte(pairPackets[0]), &left)
		json.Unmarshal([]byte(pairPackets[1]), &right)
		packets = append(packets, left, right)
		if InOrder(left, right) <= 0 {
			puzzleOne += i + 1
		}
	}

	packets = append(packets, []any{[]any{2.}}, []any{[]any{6.}})
	sort.Slice(packets, func(i, j int) bool { return InOrder(packets[i], packets[j]) < 0 })

	puzzleTwo := 1
	for i, p := range packets {
		packetStr := fmt.Sprint(p)
		if packetStr == "[[2]]" || packetStr == "[[6]]" {
			puzzleTwo *= i + 1
		}
	}
	return &Result{
		First:    puzzleOne,
		Second:   puzzleTwo,
		Duration: time.Now().Sub(begin),
	}, nil
}

// Had trouble understanding how to normalize the lists. Got help from here:
//
//	https://github.com/mnml/aoc/blob/main/2022/13/1.go
func InOrder(l, r any) int {
	lSlice, lIsSlice := l.([]any)
	rSlice, rIsSlice := r.([]any)

	switch {
	case !lIsSlice && !rIsSlice:
		return int(l.(float64) - r.(float64))
	case !lIsSlice:
		lSlice = []any{l}
	case !rIsSlice:
		rSlice = []any{r}
	}

	for i := 0; i < len(lSlice) && i < len(rSlice); i++ {
		if v := InOrder(lSlice[i], rSlice[i]); v != 0 {
			return v
		}
	}

	return len(lSlice) - len(rSlice)
}
