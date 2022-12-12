package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/echojc/aocutil"
)

func init() {
	puzzles = append(puzzles, &Day9Puzzle{})
}

type Day9Puzzle struct{}

func (day *Day9Puzzle) Number() int {
	return 9
}

func (day *Day9Puzzle) Name() string {
	return fmt.Sprintf("day_%02d", day.Number())
}

func (day *Day9Puzzle) Solution() (*Result, error) {
	begin := time.Now()
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		return nil, err
	}

	data, err := i.Strings(2022, day.Number())
	if err != nil {
		return nil, err
	}

	var (
		rope = NewRope(9)
	)
	for _, move := range data {
		dir := strings.Split(move, " ")[0]
		dist, _ := strconv.Atoi(strings.Split(move, " ")[1])
		rope.Move(dir, dist)
	}
	return &Result{
		First:    len(rope.Tail.History),
		Second:   len(rope.Tail.Last().History),
		Duration: time.Now().Sub(begin),
	}, nil
}

type Position struct {
	X, Y int
}

func (pos *Position) Move(x, y int) {
	pos.X += x
	pos.Y += y
}

func (pos *Position) Touches(other Position) bool {
	xDiff, yDiff := pos.Diff(&other)
	return xDiff >= -1 && xDiff <= 1 && yDiff >= -1 && yDiff <= 1
}

func (pos *Position) Diff(from *Position) (int, int) {
	return from.X - pos.X, from.Y - pos.Y
}

func (pos *Position) Direction(to *Position) (dir string) {
	xDiff, yDiff := pos.Diff(to)
	if yDiff > 0 {
		dir = "N"
	} else if yDiff < 0 {
		dir = "S"
	}
	if xDiff > 0 {
		dir = fmt.Sprintf("%sE", dir)
	} else if xDiff < 0 {
		dir = fmt.Sprintf("%sW", dir)
	}
	return
}

func (pos *Position) String() string {
	return fmt.Sprintf("{%d, %d}", pos.X, pos.Y)
}

type Rope struct {
	Head *Position
	Tail *Tail
}

func (r *Rope) String() string {
	return fmt.Sprintf("Head: %s; Tail: %s", r.Head, r.Tail)
}

type Tail struct {
	Position *Position
	Next     *Tail
	History  map[Position]bool
}

func (tail *Tail) Move(last *Position) {
	if tail.Position.Touches(*last) {
		return
	}
	x, y := 0, 0
	dir := []rune(tail.Position.Direction(last))
	if contains(dir, 'N') {
		y++
	} else if contains(dir, 'S') {
		y--
	}
	if contains(dir, 'E') {
		x++
	} else if contains(dir, 'W') {
		x--
	}
	tail.Position.Move(x, y)
	tail.History[*tail.Position] = true
	if tail.Next != nil {
		tail.Next.Move(tail.Position)
	}
}

func (t *Tail) Last() *Tail {
	if t.Next == nil {
		return t
	}
	return t.Next
}

func (t *Tail) String() string {
	if t.Next == nil {
		return t.Position.String()
	}
	return fmt.Sprintf("%s, %s", t.Position, t.Next)
}

func (t *Tail) Index(i int) *Tail {
	if i == 0 {
		return t
	}
	i--
	return t.Next.Index(i)
}

func (r *Rope) Move(direction string, distance int) {
	for i := 0; i < distance; i++ {
		switch direction {
		case "U":
			r.Head.Move(0, 1)
		case "D":
			r.Head.Move(0, -1)
		case "R":
			r.Head.Move(1, 0)
		case "L":
			r.Head.Move(-1, 0)
		}
		r.Tail.Move(r.Head)
	}
}

func NewRope(tailLength int) *Rope {
	return &Rope{
		Head: &Position{0, 0},
		Tail: NewTail(tailLength),
	}
}

func NewTail(length int) *Tail {
	startingPosition := Position{0, 0}
	tail := Tail{
		Position: &startingPosition,
		History:  map[Position]bool{startingPosition: true},
	}
	if length > 1 {
		tail.Next = NewTail(length - 1)
	}
	return &tail
}
