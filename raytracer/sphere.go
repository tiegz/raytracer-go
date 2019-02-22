package raytracer

import "fmt"

type Sphere struct {
	Origin    Tuple
	Radius    float64
	Transform Matrix
}

func NewSphere() Sphere { // o Tuple, r float64
	// hardcoding spheres for now
	return Sphere{NewPoint(0, 0, 0), 1, IdentityMatrix()}
}
func (s Sphere) String() string {
	return fmt.Sprintf("Sphere( %v, %.3f )", s.Origin, s.Radius)
}

func (s *Sphere) IsEqualTo(s2 Sphere) bool {
	if !s.Origin.IsEqualTo(s2.Origin) {
		return false
	} else if s.Radius != s2.Radius {
		return false
	} else if !s.Transform.IsEqualTo(s2.Transform) {
		return false
	}
	return true
}
