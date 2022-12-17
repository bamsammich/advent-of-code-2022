package main

import "fmt"

type Position struct {
	X, Y int
}

func (pos *Position) Move(x, y int) {
	pos.X += x
	pos.Y += y
}

func (pos *Position) Touches(other *Position) bool {
	xDiff, yDiff := pos.Diff(other)
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

func (pos *Position) Toward(last *Position) {
	if pos.Touches(last) {
		return
	}
	x, y := 0, 0
	dir := []rune(pos.Direction(last))
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
	pos.Move(x, y)
}

func (pos *Position) CardinalNeighbors() []Position {
	return []Position{
		{pos.X, pos.Y + 1},
		{pos.X - 1, pos.Y},
		{pos.X + 1, pos.Y},
		{pos.X, pos.Y - 1},
	}
}

func (pos *Position) DiagonalNeighbors() []Position {
	return []Position{
		{pos.X - 1, pos.Y + 1},
		{pos.X + 1, pos.Y + 1},
		{pos.X - 1, pos.Y - 1},
		{pos.X + 1, pos.Y - 1},
	}
}
