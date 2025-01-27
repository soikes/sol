package collision

import "fmt"

const quadTreeDefaultNodeCapacity = 10

type QuadTree struct {
	Boundary *Box3D

	Nodes    []*Box3D
	Capacity int

	NorthEast *QuadTree
	NorthWest *QuadTree
	SouthWest *QuadTree
	SouthEast *QuadTree
}

func NewQuadTree(boundary *Box3D, capacity int) *QuadTree {
	if capacity == 0 {
		capacity = quadTreeDefaultNodeCapacity
	}
	return &QuadTree{Boundary: boundary, Capacity: capacity}
}

func (q *QuadTree) Insert(n *Box3D) bool {
	currentNode := q

	for {
		if !currentNode.Boundary.Contains(*n) {
			return false
		}
		if len(currentNode.Nodes) < currentNode.Capacity {
			if currentNode.NorthEast == nil {
				currentNode.Nodes = append(currentNode.Nodes, n)
				return true
			}
		} else {
			if currentNode.NorthEast == nil {
				currentNode.subdivide()
			}
		}
		if currentNode.NorthEast.Boundary.Contains(*n) {
			currentNode = currentNode.NorthEast
		} else if currentNode.NorthWest.Boundary.Contains(*n) {
			currentNode = currentNode.NorthWest
		} else if currentNode.SouthWest.Boundary.Contains(*n) {
			currentNode = currentNode.SouthWest
		} else if currentNode.SouthEast.Boundary.Contains(*n) {
			currentNode = currentNode.SouthEast
		} else {
			return false
		}
	}
}

// Subdivide divides a QuadTree into four equal subdivisions.
// Nodes that are fully contained in a subdivision are moved to that subdivision.
// Nodes that span two or more subdivisions stay in the current QuadTree.
func (q *QuadTree) subdivide() {
	// Shift right to divide by 2
	midX := q.Boundary.MaxX >> 1
	midY := q.Boundary.MaxY >> 1

	ne := Box3D{MinX: midX, MaxX: q.Boundary.MaxX, MinY: midY, MaxY: q.Boundary.MaxY}
	q.NorthEast = NewQuadTree(&ne, q.Capacity)
	nw := Box3D{MinX: q.Boundary.MinX, MaxX: midX, MinY: midY, MaxY: q.Boundary.MaxY}
	q.NorthWest = NewQuadTree(&nw, q.Capacity)
	sw := Box3D{MinX: q.Boundary.MinX, MaxX: midX, MinY: q.Boundary.MinY, MaxY: midY}
	q.SouthWest = NewQuadTree(&sw, q.Capacity)
	se := Box3D{MinX: midX, MaxX: q.Boundary.MaxX, MinY: q.Boundary.MinY, MaxY: midY}
	q.SouthEast = NewQuadTree(&se, q.Capacity)

	var i = 0
	for _, n := range q.Nodes {
		if q.NorthEast.Insert(n) {
			continue
		} else if q.NorthWest.Insert(n) {
			continue
		} else if q.SouthWest.Insert(n) {
			continue
		} else if q.SouthEast.Insert(n) {
			continue
		} else {
			q.Nodes[i] = n
			i++
		}
	}
	for j := i; j < len(q.Nodes); j++ {
		q.Nodes[j] = nil
	}
	q.Nodes = q.Nodes[:i]
}

func (q *QuadTree) Neighbours(node *Box3D) []*Box3D {
	var stack []*QuadTree
	stack = append(stack, q)

	for len(stack) > 0 {
		currentNode := stack[len(stack)-1]
		stack[len(stack)-1] = nil
		stack = stack[:len(stack)-1]

		for _, n := range currentNode.Nodes {
			if n == node {
				return currentNode.Nodes
			}
		}

		if currentNode.NorthEast != nil {
			stack = append(stack, currentNode.SouthEast, currentNode.SouthWest, currentNode.NorthWest, currentNode.NorthEast)
		}
	}

	return nil
}

const spacer = "+--- "
const indent = "|\t"

func (q *QuadTree) String() string {
	var out string
	var stack []*QuadTree
	var depthMap = make(map[*QuadTree]int)

	stack = append(stack, q)
	depthMap[q] = 0

	for len(stack) > 0 {
		currentNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		currentDepth := 0

		if depth, ok := depthMap[currentNode]; ok {
			for i := 0; i < depth; i++ {
				out += indent
			}
			out += spacer
			out += fmt.Sprintf("%s\t| ", currentNode.Boundary)
			for _, n := range currentNode.Nodes {
				out += fmt.Sprintf("%s, ", n.String())
			}
			out += "\n"
			currentDepth = depth
		}

		if currentNode.NorthEast != nil {
			stack = append(stack, currentNode.SouthEast, currentNode.SouthWest, currentNode.NorthWest, currentNode.NorthEast)
			depthMap[currentNode.SouthEast] = currentDepth + 1
			depthMap[currentNode.SouthWest] = currentDepth + 1
			depthMap[currentNode.NorthWest] = currentDepth + 1
			depthMap[currentNode.NorthEast] = currentDepth + 1
		}
	}

	return out
}
