package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day10Puzzle{})
}

type Day10Puzzle struct{}

func (day *Day10Puzzle) Number() int {
	return 10
}

func (day *Day10Puzzle) Name() string {
	return fmt.Sprintf("day_%02d", day.Number())
}

func (day *Day10Puzzle) Solution() (*Result, error) {
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
		register      = 1
		cycle         = 0
		sumCheckpoint = 20
		rowLength     = 40
		puzzleOne     []int
		puzzleTwo     bytes.Buffer
	)
	for _, instruction := range data {
		var val int

		for i := 0; i < 2; i++ {
			drawPixel(cycle%rowLength, register, &puzzleTwo)
			cycle++
			if instruction == "noop" {
				break
			}
			val, _ = strconv.Atoi(strings.Split(instruction, " ")[1])
			if cycle%sumCheckpoint == 0 {
				puzzleOne = append(puzzleOne, cycle*register)
				sumCheckpoint += 40
			}
		}
		register += val
	}

	return &Result{
		First:    sumInts(puzzleOne),
		Second:   drawCRT(puzzleTwo.String(), rowLength),
		Duration: time.Now().Sub(begin),
	}, nil
}

func drawPixel(cycle, register int, stream *bytes.Buffer) {
	if cycle >= register-1 && cycle <= register+1 {
		stream.WriteString("#")
	} else {
		stream.WriteString(".")
	}
	return
}

func drawCRT(stream string, rowLength int) (out []string) {
	for len(stream) > rowLength-1 {
		out = append(out, stream[:rowLength])
		stream = stream[rowLength:]
	}
	if len(stream) > 0 {
		out = append(out, stream)
	}
	return
}
