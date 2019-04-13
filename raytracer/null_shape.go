package raytracer

import "fmt"

type NullShape struct {
}

func NewNullShape() Shape {
	return NewShape(NullShape{})
}

/////////////////////////
// ShapeInterface methods
/////////////////////////

// TODO can we remove Shape arg somehow? It's only there because ShapeInterface
// has no knowledge of its parent, but we need to put its aprent in the Intersection :(
func (ns NullShape) LocalIntersect(r Ray, shape *Shape) Intersections {
	return Intersections{}
}

func (ns NullShape) LocalNormalAt(worldPoint Tuple) Tuple {
	return NewVector(worldPoint.X, worldPoint.Y, worldPoint.Z)
}

func (ns NullShape) localIsEqualTo(s2 ShapeInterface) bool {
	// TODO
	return true
}

func (ns NullShape) String() string {
	return fmt.Sprintf("TestShape( NULL )")
}

func (ns NullShape) localString() string {
	return ns.String()
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (ns NullShape) localType() string {
	return "NullShape"
}
