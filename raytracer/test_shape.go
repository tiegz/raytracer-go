package raytracer

import "fmt"

// This is just a dummy "child" of Shape, for testing purposes.
type TestShape struct {
}

func NewTestShape() *Shape {
	return NewShape(TestShape{})
}

/////////////////////////
// ShapeInterface methods
/////////////////////////

func (ts TestShape) LocalBounds() BoundingBox {
	return NewBoundingBox(NewPoint(-1, -1, -1), NewPoint(1, 1, 1))
}

func (ts TestShape) LocalIntersect(r Ray, shape *Shape) Intersections {
	shape.SavedRay = &r
	return Intersections{}
}

func (ts TestShape) LocalNormalAt(localPoint Tuple, hit *Intersection) Tuple {
	return NewVector(localPoint.X, localPoint.Y, localPoint.Z)
}

func (s TestShape) localIsEqualTo(s2 ShapeInterface) bool {
	// TODO: when we are able to move SavedRay from Shape to TestShape, uncomment this
	// s2TestShape := s2.(TestShape)
	// if !s.SavedRay.IsEqualTo(s2TestShape.SavedRay) {
	// 	return false
	// }
	return true
}

func (s TestShape) String() string {
	return fmt.Sprintf("TestShape( )")
}

func (s TestShape) localString() string {
	return s.String()
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (s TestShape) localType() string {
	return "TestShape"
}
