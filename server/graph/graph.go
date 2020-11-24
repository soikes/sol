package graph

import (
	"fmt"
)

type SearchMethod int

const (
	MethodDFS SearchMethod = iota
	MethodBFS
)

type Graph struct {
	vertices []*Vertex
}

func (g *Graph) Count() int {
	return len(g.vertices)
}

func (g *Graph) Add(v *Vertex) {
	fmt.Println(`adding vertex: `, v)
	g.vertices = append(g.vertices, v)
}

func (g *Graph) Root() *Vertex {
	if g.Count() != 0 {
		return g.vertices[0]
	}
	return nil
}

func Link(v1, v2 *Vertex, cost int) {
	v1.Edges = append(v1.Edges, &Edge{V: v2, Cost: cost})
	v2.Edges = append(v2.Edges, &Edge{V: v1, Cost: cost})
}

func IsAdjacent(v1 *Vertex, v2 *Vertex) bool {
	for _, e := range v1.Edges {
		if e.V == v2 {
			return true
		}
	}
	return false
}

type Edge struct {
	V *Vertex
	Cost int
}

type Vertex struct {
	Edges []*Edge
	Value interface{}
}

func (v *Vertex) Neighbours() []*Vertex {
	n := []*Vertex{}
	for _, e := range v.Edges {
		n = append(n, e.V)
	}
	return n
}

func (g *Graph) Search(val interface{}, method SearchMethod) *Vertex {
	switch method {
	case MethodDFS:
		return g.searchDFS(val, nil, []*Vertex{})
	case MethodBFS:
		return g.searchBFS(val, nil, []*Vertex{})
	default:
		return g.searchDFS(val, nil, []*Vertex{})
	}
}

func (g *Graph) searchDFS(val interface{}, cur *Vertex, visited []*Vertex) *Vertex {
	if len(g.vertices) == 0 {
		return nil //TODO probably not a good decision
	}
	if cur == nil {
		cur = g.Root()
	}
	fmt.Println(`cur:`, cur.Value)
	if cur.Value == val {
		return cur
	}

	fmt.Println(`visited`, cur.Value)
	visited = append(visited, cur)

	for _, e := range cur.Edges {
		fmt.Println(`edge to:`, e.V.Value)
		vis := false
		for _, v := range visited {
			if e.V == v {
				fmt.Println(`already visited, returning`)
				vis = true
			}
		}
		fmt.Println(`going deeper`)
		if !vis {
			return g.searchDFS(val, e.V, visited)
		}
	}
	fmt.Println(`returning from function`)
	return nil
}

func (g *Graph) searchBFS(value interface{}, root *Vertex, visited []*Vertex) *Vertex {
	return nil
}

func (g *Graph) ShortestPath(value interface{}) []*Vertex {
	return nil
}

func (g *Graph) ToString() string {
	out := ``
	for _, v := range g.vertices {
		out += fmt.Sprintf(`vertex: %s, links: `, v.Value)
		for _, n := range v.Neighbours() {
			out += fmt.Sprintf(`%s `, n.Value)
		}
		out += fmt.Sprintln()
	}
	return out
}