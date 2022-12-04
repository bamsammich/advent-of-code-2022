package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day4Puzzle{})
}

type Day4Puzzle struct{}

func (day *Day4Puzzle) Name() string {
	return "day_4"
}

type Section struct {
	Start, End int
}

func (s *Section) Contains(c *Section) bool {
	if s.Start <= c.Start && s.End >= c.End {
		return true
	}
	return false
}

func (s *Section) Overlaps(c *Section) bool {
	if s.End <= c.End && s.End >= c.Start {
		return true
	}
	if s.End >= c.End && s.Start <= c.End {
		return true
	}
	return false
}

func SectionFromElf(elf string) *Section {
	sectionStr := strings.Split(elf, "-")
	start, _ := strconv.Atoi(sectionStr[0])
	end, _ := strconv.Atoi(sectionStr[1])
	return &Section{start, end}
}

func (day *Day4Puzzle) Solution() (*Result, error) {
	begin := time.Now()
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	data, err := i.Strings(2022, 4)
	if err != nil {
		return nil, err
	}

	var countContains = 0
	var countOverlaps = 0
	for _, pair := range data {
		elves := strings.Split(pair, ",")
		elfOne := SectionFromElf(elves[0])
		elfTwo := SectionFromElf(elves[1])
		if elfOne.Contains(elfTwo) || elfTwo.Contains(elfOne) {
			countContains++
		}
		if elfOne.Overlaps(elfTwo) {
			countOverlaps++
		}
	}

	return &Result{
		First:    countContains,
		Second:   countOverlaps,
		Duration: time.Now().Sub(begin),
	}, nil
}
