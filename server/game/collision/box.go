package collision

import "fmt"

// TODO Does this need to be an interface? e.g. If I want to add nodes to a QuadTree
// that have custom properties but can still be used as a Box3D
type Box3D struct {
	MinX int
	MaxX int
	MinY int
	MaxY int
	MinZ int
	MaxZ int
}

func (a *Box3D) Intersects(b Box3D) bool {
	return a.MinX <= b.MaxX &&
		a.MaxX >= b.MinX &&
		a.MinY <= b.MaxY &&
		a.MaxY >= b.MinY &&
		a.MinZ <= b.MaxZ &&
		a.MaxZ >= b.MinZ
}

func (a *Box3D) Contains(b Box3D) bool {
	return a.MinX <= b.MinX &&
		a.MaxX >= b.MaxX &&
		a.MinY <= b.MinY &&
		a.MaxY >= b.MaxY &&
		a.MinZ <= b.MinZ &&
		a.MaxZ >= b.MaxZ
}

func (a *Box3D) String() string {
	return fmt.Sprintf("{(%d,%d),(%d,%d),(%d,%d)}", a.MinX, a.MaxX, a.MinY, a.MaxY, a.MinZ, a.MaxZ)
}
