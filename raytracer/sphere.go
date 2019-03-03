package raytracer

import "fmt"

type Sphere struct {
	Origin    Tuple
	Radius    float64
	Transform Matrix
	Material  Material
}

func NewSphere() Sphere { // o Tuple, r float64
	// hardcoding spheres for now
	return Sphere{NewPoint(0, 0, 0), 1, IdentityMatrix(), DefaultMaterial()}
}
func (s Sphere) String() string {
	return fmt.Sprintf("Sphere( %v, %.1f )", s.Origin, s.Radius)
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

// Given a point, return the vector from origin of the sphere (the "normal").
func (s *Sphere) NormalAt(worldPoint Tuple) Tuple {
	// convert the "world-space" point to "object-space"
	inverseTransform := s.Transform.Inverse()
	objectPoint := inverseTransform.MultiplyByTuple(worldPoint)
	objectNormal := objectPoint.Subtract(s.Origin)

	// convert the "object-space" normal (vector) to "world-space"
	inverseTransformTransposed := inverseTransform.Transpose()
	worldNormal := inverseTransformTransposed.MultiplyByTuple(objectNormal)

	// HACK " ... should be finding submatrix(transform, 3, 3) first, and multiplying by the inverse and trans- pose of that."
	worldNormal.W = 0

	return worldNormal.Normalized()
}
