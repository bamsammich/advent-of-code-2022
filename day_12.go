package main

import (
	"container/heap"
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
		graph         = NewGraph()
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

// Dijkstra solution found at https://dev.to/douglasmakey/implementation-of-dijkstra-using-heap-in-go-6e3
type Path struct {
	Value int
	Nodes []Position
}

type ShortestPath []Path

func (sp ShortestPath) Len() int { return len(sp) }

func (sp ShortestPath) Less(i, j int) bool { return sp[i].Value < sp[j].Value }

func (sp ShortestPath) Swap(i, j int) { sp[i], sp[j] = sp[j], sp[i] }

func (sp *ShortestPath) Push(x any) {
	*sp = append(*sp, x.(Path))
}

func (sp *ShortestPath) Pop() any {
	old := *sp
	*sp = old[0 : len(old)-1]
	return old[len(old)-1]
}

type PathHeap struct {
	Values *ShortestPath
}

func NewPathHeap() *PathHeap {
	return &PathHeap{Values: &ShortestPath{}}
}

func (ph *PathHeap) Push(p Path) {
	heap.Push(ph.Values, p)
}

func (ph *PathHeap) Pop() Path {
	p := heap.Pop(ph.Values)
	return p.(Path)
}

type Edge struct {
	Node   Position
	Weight int
}

type Graph struct {
	Nodes map[Position][]Edge
}

func NewGraph() *Graph {
	return &Graph{Nodes: make(map[Position][]Edge)}
}

func (g *Graph) AddEdge(src, dest Position, weight int) {
	g.Nodes[src] = append(g.Nodes[src], Edge{Node: dest, Weight: weight})
	// g.Nodes[dest] = append(g.Nodes[dest], Edge{Node: src, Weight: weight})
}

func (g *Graph) GetEdges(node Position) []Edge {
	return g.Nodes[node]
}

func (g *Graph) GetPath(src, dest Position) (int, []Position) {
	h := NewPathHeap()
	h.Push(Path{Value: 0, Nodes: []Position{src}})
	visited := make(map[Position]bool)

	for len(*h.Values) > 0 {
		p := h.Pop()
		node := p.Nodes[len(p.Nodes)-1]

		if visited[node] {
			continue
		}
		if node == dest {
			return p.Value, p.Nodes
		}

		for _, e := range g.GetEdges(node) {
			if !visited[e.Node] {
				// We calculate the total spent so far plus the cost and the path of getting here
				h.Push(Path{
					Value: p.Value + e.Weight,
					Nodes: append([]Position{}, append(p.Nodes, e.Node)...),
				})
			}
		}
		visited[node] = true
	}
	return 0, nil
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
