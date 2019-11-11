package raytracer

import (
	"fmt"
)

type ShapeInterface interface {
	LocalNormalAt(Tuple, Intersection) Tuple
	LocalIntersect(Ray, *Shape) Intersections
	LocalBounds() BoundingBox
	localIsEqualTo(ShapeInterface) bool
	localType() string
	localString() string
}

// Shape is a general shape (Transform+Material), with the specific type of shape stored as a ShapeInterface in LocalShape.
// LocaleShape-specific functions are prefixed with "local", e.g. localFoo().
type Shape struct {
	LocalShape ShapeInterface // not using anonymous embedded field mostly bc of IsEqualTo()... we have to pass the LocalShape, not the Shape
	Transform  Matrix
	Material   Material
	SavedRay   Ray // TODO replace this later, it's only for testing purposes with TestShape
	Parent     *Shape
	Label      string
	Shadows    bool
}

func NewShape(si ShapeInterface) Shape {
	return Shape{
		LocalShape: si,
		Transform:  IdentityMatrix(),
		Material:   DefaultMaterial(),
		SavedRay:   NullRay(),
		Shadows:    true, // does this shape cast shadows?
	}
}

func (s *Shape) Intersect(r Ray) Intersections {
	// Instead of applying object's transformation to object, we can just apply
	// the inverse of the transformation to the ray.
	r = r.Transform(s.Transform.Inverse())
	return s.LocalShape.LocalIntersect(r, s)
}

func (s *Shape) NormalAt(worldPoint Tuple, i Intersection) Tuple {
	objectPoint := s.WorldToObject(worldPoint)
	objectNormal := s.LocalShape.LocalNormalAt(objectPoint, i)
	return s.NormalToWorld(objectNormal)
}

// Transforms a point in world space to object space, accounting for the chain of parents in between.
func (s *Shape) WorldToObject(worldPoint Tuple) Tuple {
	if s.Parent != nil {
		worldPoint = s.Parent.WorldToObject(worldPoint)
	}
	inverseTransform := s.Transform.Inverse()
	return inverseTransform.MultiplyByTuple(worldPoint)
}

// Transforms a vector in object space to world space, accounting for the chain of parents in between.
func (s *Shape) NormalToWorld(normal Tuple) Tuple {
	inverseTransform := s.Transform.Inverse()
	tranposedInverseTranform := inverseTransform.Transpose()

	normal = tranposedInverseTranform.MultiplyByTuple(normal)
	normal.W = 0 // HACK: " ... should be finding submatrix(transform, 3, 3) first, and multiplying by the inverse and trans- pose of that."
	normal = normal.Normalized()

	if s.Parent != nil {
		normal = s.Parent.NormalToWorld(normal)
	}

	return normal
}

func (s *Shape) AddChildren(shapes ...*Shape) {
	// HACK: sucks to have to create new Group. When we test out replacing value LocalShape with *LocalShape, we can just alter the Group directly? Otherwise, generalize this into a Copy() method for at least Group.
	if s.LocalShape.localType() == "Group" {
		group := s.LocalShape.(Group)
		group.Children = append(group.Children, shapes...)
		for _, shape := range group.Children {
			shape.Parent = s
		}
		s.LocalShape = group
	} else {
		// TODO: return error
		panic("Should never AddChildren() to non-Group!\n")
	}
}

// NB: this returns true for "regular shape includes itself"
func (s *Shape) Includes(s2 *Shape) bool {
	if s.IsEqualTo(*s2) {
		return true
	}

	switch localShape := s.LocalShape.(type) {
	case Group:
		for _, child := range localShape.Children {
			if child.Includes(s2) {
				return true
			}
		}
	case Csg:
		if localShape.Left.Includes(s2) || localShape.Right.Includes(s2) {
			return true
		}
	}

	return false
}

func (s Shape) Bounds() BoundingBox {
	return s.LocalShape.LocalBounds()
}

func (s Shape) ParentSpaceBounds() BoundingBox {
	b := s.Bounds()
	return b.Transform(s.Transform)
}

func (s Shape) String() string {
	// return fmt.Sprintf("Shape( \n  %v \n  %v \n  %v\n)", s.Material, s.Transform, s.LocalShape.localString())
	return fmt.Sprintf("Shape( Label: %v LocalShape: %v )", s.Label, s.LocalShape.localString())
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
