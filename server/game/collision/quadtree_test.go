package collision

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQuadTree(t *testing.T) {
	boundary := Box3D{MinX: 0, MaxX: 10, MinY: 0, MaxY: 10, MinZ: 0, MaxZ: 0}
	a := Box3D{MinX: 0, MaxX: 4, MinY: 0, MaxY: 4, MinZ: 0, MaxZ: 0}
	b := Box3D{MinX: 2, MaxX: 6, MinY: 2, MaxY: 6, MinZ: 0, MaxZ: 0}
	c := Box3D{MinX: 1, MaxX: 2, MinY: 1, MaxY: 2, MinZ: 0, MaxZ: 0}
	d := Box3D{MinX: 5, MaxX: 6, MinY: 5, MaxY: 6, MinZ: 0, MaxZ: 0}
	e := Box3D{MinX: 1, MaxX: 2, MinY: 1, MaxY: 2, MinZ: 0, MaxZ: 0}
	tree := NewQuadTree(&boundary, 2)

	tree.Insert(&a)
	n1 := tree.Neighbours(&a)
	if len(n1) != 1 {
		t.Errorf("neighbours: want 1, got %s", n1)
	}

	tree.Insert(&b)
	n2 := tree.Neighbours(&b)
	if len(n2) != 2 {
		t.Errorf("neighbours: want 2, got %s", n2)
	}

	tree.Insert(&c) // Should cause a subdivide
	n3 := tree.Neighbours(&c)
	if len(n3) != 2 {
		t.Errorf("neighbours: want 2, got %s", n3)
	}
	n4 := tree.Neighbours(&b)
	if len(n4) != 1 {
		t.Errorf("neighbours: want 1, got %s", n4)
	}

	tree.Insert(&d)
	tree.Insert(&e)
	fmt.Println(tree)
}

var benchTree *QuadTree
var searchNode *Box3D

func setupBenchmarks(capacity int) {
	boundary := Box3D{MinX: 0, MaxX: 100, MinY: 0, MaxY: 100, MinZ: 0, MaxZ: 0}
	benchTree = NewQuadTree(&boundary, capacity)
	searchNode = &Box3D{MinX: 98, MaxX: 99, MinY: 98, MaxY: 99, MinZ: 0, MaxZ: 0}
	for i := 0; i < 10000; i++ {
		a := rand.Intn(100)
		b := rand.Intn(100)
		c := rand.Intn(100)
		d := rand.Intn(100)
		n := Box3D{}
		if a < b {
			n.MinX = a
			n.MaxX = b
		} else {
			n.MinX = b
			n.MaxX = a
		}
		if c < d {
			n.MinY = c
			n.MaxY = d
		} else {
			n.MinY = d
			n.MaxY = c
		}
		benchTree.Insert(&n)
	}
}

func BenchmarkQuadTree(b *testing.B) {
	b.Run("Capacity5", func(b *testing.B) {
		setupBenchmarks(5)
		for i := 0; i < b.N; i++ {
			_ = benchTree.Neighbours(searchNode)
		}
	})
	b.Run("Capacity10", func(b *testing.B) {
		setupBenchmarks(10)
		for i := 0; i < b.N; i++ {
			_ = benchTree.Neighbours(searchNode)
		}
	})
	b.Run("Capacity15", func(b *testing.B) {
		setupBenchmarks(15)
		for i := 0; i < b.N; i++ {
			_ = benchTree.Neighbours(searchNode)
		}
	})
	b.Run("Capacity100", func(b *testing.B) {
		setupBenchmarks(100)
		for i := 0; i < b.N; i++ {
			_ = benchTree.Neighbours(searchNode)
		}
	})
}
