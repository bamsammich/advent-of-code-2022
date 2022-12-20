package main

import "container/heap"

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

type NodeGraph struct {
	Nodes map[Position][]Edge
}

func NewNodeGraph() *NodeGraph {
	return &NodeGraph{Nodes: make(map[Position][]Edge)}
}

func (g *NodeGraph) AddEdge(src, dest Position, weight int) {
	g.Nodes[src] = append(g.Nodes[src], Edge{Node: dest, Weight: weight})
	// g.Nodes[dest] = append(g.Nodes[dest], Edge{Node: src, Weight: weight})
}

func (g *NodeGraph) GetEdges(node Position) []Edge {
	return g.Nodes[node]
}

func (g *NodeGraph) GetPath(src, dest Position) (int, []Position) {
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
