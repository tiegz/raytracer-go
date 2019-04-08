package raytracer

import "fmt"

// This is just a dummy "child" of Shape, for testing purposes.
type TestShape struct {
	SavedRay Ray
}

func NewTestShape() Shape {
	return NewShape(TestShape{})
}

/////////////////////////
// ShapeInterface methods
/////////////////////////

func (ts TestShape) localIntersect(r Ray, shape *Shape) Intersections {
	shape.SavedRay = r // the book has this on the TestShape object, but for Go I've put it on Shape :shruggie:
	return Intersections{}
}

func (ts TestShape) localNormalAt(worldPoint Tuple) Tuple {
	return NewVector(worldPoint.X, worldPoint.Y, worldPoint.Z)
}

func (s TestShape) localIsEqualTo(s2 ShapeInterface) bool {
	s2TestShape := s2.(TestShape)
	if !s.SavedRay.IsEqualTo(s2TestShape.SavedRay) {
		return false
	}
	return true
}

func (s TestShape) String() string {
	return fmt.Sprintf("TestShape( %v )", s.SavedRay)
}

func (s TestShape) localString() string {
	return s.String()
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (s TestShape) localType() string {
	return "TestShape"
}
