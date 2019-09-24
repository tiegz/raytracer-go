package raytracer

import (
	"fmt"
	"math"
)

type Cube struct {
	Origin Tuple
}

func NewCube() Shape {
	return NewShape(&Cube{NewPoint(0, 0, 0)})
}

func (c Cube) String() string {
	return fmt.Sprintf("Cube( %v )", c.Origin)
}

/////////////////////////
// ShapeInterface methods
/////////////////////////

func (c Cube) localString() string {
	return c.String()
}

// TODO can we remove Shape arg somehow? It's only there because ShapeInterface
// has no knowledge of its parent, but we need to put its aprent in the Intersection :(
// We treat a cube like 6 planes, with 2 parallel planes per axis.
func (c Cube) LocalIntersect(r Ray, shape *Shape) Intersections {
	xTMin, xTMax := c.checkAxis(r.Origin.X, r.Direction.X)
	yTMin, yTMax := c.checkAxis(r.Origin.Y, r.Direction.Y)
	zTMin, zTMax := c.checkAxis(r.Origin.Z, r.Direction.Z)

	// ... The intersection of the ray with that square will always be those two points: the largest minimum t value and the smallest maximum t value. ...
	tMin := maxFloat64(xTMin, yTMin, zTMin)
	tMax := minFloat64(xTMax, yTMax, zTMax)

	if tMin > tMax {
		return Intersections{}
	}

	return Intersections{
		Intersection{tMin, *shape},
		Intersection{tMax, *shape},
	}
}

func (c Cube) checkAxis(origin float64, direction float64) (float64, float64) {
	tMinNumerator := (-1 - origin)
	tMaxNumerator := (1 - origin)
	var tMin, tMax float64

	if math.Abs(direction) >= EPSILON {
		tMin = tMinNumerator / direction
		tMax = tMaxNumerator / direction
	} else {
		tMin = tMinNumerator * math.Inf(1)
		tMax = tMaxNumerator * math.Inf(1)
	}
	if tMin > tMax {
		return tMax, tMin
	} else {
		return tMin, tMax
	}
}

func (c Cube) LocalNormalAt(localPoint Tuple) Tuple {
	maxC := maxFloat64(math.Abs(localPoint.X), math.Abs(localPoint.Y), math.Abs(localPoint.Z))
	if maxC == math.Abs(localPoint.X) {
		return NewVector(localPoint.X, 0, 0)
	} else if maxC == math.Abs(localPoint.Y) {
		return NewVector(0, localPoint.Y, 0)
	} else {
		return NewVector(0, 0, localPoint.Z)
	}
}

func (c Cube) localIsEqualTo(c2 ShapeInterface) bool {
	// NB I still don't know why we need to do a type assertion to *Cube
	// instead of Cube, but it fixes this panic when comparing two
	// cubes: 'panic: interface conversion: raytracer.ShapeInterface is *raytracer.Cube, not raytracer.Cube [recovered]
	c2Cube := c2.(*Cube)
	if !c.Origin.IsEqualTo(c2Cube.Origin) {
		return false
	}
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (c Cube) localType() string {
	return "Cube"
}
