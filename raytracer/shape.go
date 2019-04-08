package raytracer

import (
	"fmt"
)

type ShapeInterface interface {
	localNormalAt(Tuple) Tuple
	localIntersect(Ray, *Shape) Intersections
	localIsEqualTo(ShapeInterface) bool
	localType() string
	localString() string
}

// Shape is a general shape (Transform+Material), with the specific type of shape stored as a ShapeInterface in LocalShape.
type Shape struct {
	LocalShape ShapeInterface // not using anonymous embedded field mostly bc of IsEqualTo()... we have to pass the LocalShape, not the Shape
	Transform  Matrix
	Material   Material
	SavedRay   Ray // TODO replace this later, it's only for testing purposes with TestShape
}

func NewShape(si ShapeInterface) Shape {
	return Shape{LocalShape: si, Transform: IdentityMatrix(), Material: DefaultMaterial()}
}

func (s *Shape) Intersect(r Ray) Intersections {
	// Instead of applying object's transformation to object, we can just apply
	// the inverse of the transformation to the ray.
	localRay := r.Transform(s.Transform.Inverse())
	return s.LocalShape.localIntersect(localRay, s)
}

func (s *Shape) NormalAt(p Tuple) Tuple {
	// convert the "world-space" point to "object-space"
	inverseTransform := s.Transform.Inverse()
	localPoint := inverseTransform.MultiplyByTuple(p)
	localNormal := s.LocalShape.localNormalAt(localPoint)

	// convert the "object-space" normal (vector) to "world-space"
	inverseTransformTransposed := inverseTransform.Transpose()
	worldNormal := inverseTransformTransposed.MultiplyByTuple(localNormal)

	// HACK " ... should be finding submatrix(transform, 3, 3) first, and multiplying by the inverse and trans- pose of that."
	worldNormal.W = 0

	return worldNormal.Normalized()
}

func (s Shape) String() string {
	return fmt.Sprintf("Shape( \n  %v \n  %v \n  %v\n)", s.Transform, s.Material, s.LocalShape.localString())
}

func (s *Shape) IsEqualTo(s2 Shape) bool {
	st1 := s.LocalShape.localType()
	st2 := s2.LocalShape.localType()
	if st1 != st2 {
		return false
	} else if !s.Material.IsEqualTo(s2.Material) {
		return false
	} else if !s.Transform.IsEqualTo(s2.Transform) {
		return false
	} else {
		return s.LocalShape.localIsEqualTo(s2.LocalShape)
	}
}
