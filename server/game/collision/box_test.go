package collision

import "testing"

func TestIntersects(t *testing.T) {
	a := Box3D{MinX: 0, MaxX: 4, MinY: 0, MaxY: 4, MinZ: 0, MaxZ: 4}
	b := Box3D{MinX: 2, MaxX: 6, MinY: 2, MaxY: 6, MinZ: 2, MaxZ: 6}
	c := Box3D{MinX: 1, MaxX: 2, MinY: 1, MaxY: 2, MinZ: 1, MaxZ: 2}
	d := Box3D{MinX: 5, MaxX: 6, MinY: 5, MaxY: 6, MinZ: 5, MaxZ: 6}

	if intersects := a.Intersects(b); !intersects {
		t.Errorf("intersects: want: %t, got: %t\n", true, intersects)
	}

	if intersects := a.Intersects(c); !intersects {
		t.Errorf("intersects: want: %t, got: %t\n", true, intersects)
	}

	if intersects := a.Intersects(d); intersects {
		t.Errorf("intersects: want: %t, got: %t\n", false, intersects)
	}
}

func TestContains(t *testing.T) {
	a := Box3D{MinX: 0, MaxX: 4, MinY: 0, MaxY: 4, MinZ: 0, MaxZ: 4}
	b := Box3D{MinX: 2, MaxX: 6, MinY: 2, MaxY: 6, MinZ: 2, MaxZ: 6}
	c := Box3D{MinX: 1, MaxX: 2, MinY: 1, MaxY: 2, MinZ: 1, MaxZ: 2}
	// d := Box3D{MinX: 0, MaxX: 10, MinY: 0, MaxY: 10, MinZ: 0, MaxZ: 0}
	// e := Box3D{MinX: 0, MaxX: 5, MinY: 0, MaxY: 5}

	if contains := a.Contains(b); contains {
		t.Errorf("contains: want: %t, got: %t\n", false, contains)
	}

	if contains := a.Contains(c); !contains {
		t.Errorf("contains: want: %t, got: %t\n", true, contains)
	}
}
