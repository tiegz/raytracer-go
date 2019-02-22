package raytracer

import "fmt"

type Sphere struct {
	Origin Tuple
	Radius float64
}

func NewSphere() Sphere { // o Tuple, r float64
	// hardcoding spheres for now
	return Sphere{NewPoint(0, 0, 0), 1.0}
}
func (s Sphere) String() string {
	return fmt.Sprintf("Sphere( %v, %.3f )", s.Origin, s.Radius)
}
