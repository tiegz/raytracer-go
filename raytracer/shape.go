package raytracer

import (
	"fmt"
)

type ShapeInterface interface {
	LocalNormalAt(Tuple, *Intersection) Tuple
	LocalIntersect(Ray, *Shape) Intersections
	LocalBounds() BoundingBox
	localIsEqualTo(ShapeInterface) bool
	localType() string
	localString() string
}

// Shape is a general shape (Transform+Material), with the specific type of shape stored as a ShapeInterface in LocalShape.
// LocaleShape-specific functions are prefixed with "local", e.g. localFoo().
type Shape struct {
	Label            string
	LocalShape       ShapeInterface // not using anonymous embedded field mostly bc of IsEqualTo()... we have to pass the LocalShape, not the Shape
	Transform        Matrix         // WARNING: don't set Transform directly, use SetTransform()
	InverseTransform Matrix
	Material         *Material
	SavedRay         *Ray // TODO replace this later, it's only for testing purposes with TestShape
	Parent           *Shape
	Shadows          bool
}

func NewShape(si ShapeInterface) *Shape {
	s := Shape{
		LocalShape: si,
		Material:   DefaultMaterial(),
		Shadows:    true, // does this shape cast shadows?
	}
	s.SetTransform(IdentityMatrix())
	return &s
}

func (s *Shape) SetTransform(m Matrix) {
	s.Transform = m
	s.InverseTransform = m.Inverse()
}

func (s *Shape) Intersect(r Ray) Intersections {
	// Instead of applying object's transformation to object, we can just apply
	// the inverse of the transformation to the ray.
	r = r.Transform(s.InverseTransform)
	return s.LocalShape.LocalIntersect(r, s)
}

func (s *Shape) NormalAt(worldPoint Tuple, i *Intersection) Tuple {
	objectPoint := s.WorldToObject(worldPoint)
	objectNormal := s.LocalShape.LocalNormalAt(objectPoint, i)
	return s.NormalToWorld(objectNormal)
}

// Transforms a point in world space to object space, accounting for the chain of parents in between.
func (s *Shape) WorldToObject(worldPoint Tuple) Tuple {
	if s.Parent != nil {
		worldPoint = s.Parent.WorldToObject(worldPoint)
	}
	return s.InverseTransform.MultiplyByTuple(worldPoint)
}

// Transforms a vector in object space to world space, accounting for the chain of parents in between.
func (s *Shape) NormalToWorld(normal Tuple) Tuple {
	tranposedInverseTranform := s.InverseTransform.Transpose()

	normal = tranposedInverseTranform.MultiplyByTuple(normal)
	normal.W = 0 // HACK: " ... should be finding submatrix(transform, 3, 3) first, and multiplying by the inverse and trans- pose of that."
	normal = normal.Normalized()

	if s.Parent != nil {
		normal = s.Parent.NormalToWorld(normal)
	}

	return normal
}

// If the shape is a Group and has at least +threshhold+ children,
// divide the children into new subgroups, to create the BVH.
func (s *Shape) Divide(threshhold int) {
	switch localShape := s.LocalShape.(type) {
	case Group:
		if threshhold <= len(localShape.Children) {
			l, r := s.PartitionChildren()
			if len(l) > 0 {
				s.MakeSubGroup(l...)
			}
			if len(r) > 0 {
				s.MakeSubGroup(r...)
			}
		}
		for _, child := range s.LocalShape.(Group).Children {
			child.Divide(threshhold)
		}
	case Csg:
		s.LocalShape.(Csg).Left.Divide(threshhold)
		s.LocalShape.(Csg).Right.Divide(threshhold)
	default:
		// no-op
	}
}

// Adds a child that is a Group containing the given Shapes.
func (s *Shape) MakeSubGroup(shapes ...*Shape) {
	if s.LocalShape.localType() != "Group" {
		// TODO: can we move this to just Group logic instead of all Shape logic?
		// TODO: return error
		panic("Should never AddChildren() to non-Group!\n")
	}

	subGroup := NewGroup()
	subGroup.Label = "NewSubGroup"
	subGroup.AddChildren(shapes...)
	s.AddChildren(subGroup)
}

// Adds one or more children to this Group.
func (s *Shape) AddChildren(shapes ...*Shape) {
	if s.LocalShape.localType() != "Group" {
		// TODO: can we move this to just Group logic instead of all Shape logic?
		// TODO: return error
		panic("Should never AddChildren() to non-Group!\n")
	}

	// HACK: sucks to have to create new Group. When we test out replacing value LocalShape with *LocalShape, we can just alter the Group directly? Otherwise, generalize this into a Copy() method for at least Group.
	group := s.LocalShape.(Group)
	group.Children = append(group.Children, shapes...)
	for _, shape := range group.Children {
		shape.Parent = s
	}
	s.LocalShape = group
}

// TODO: can we move this to just Group logic instead of all Shape logic?
// Divides the children into 2 subgroups, based on which sub-bounding-box
// they are located in. Any shapes that are located in both remain in this Group.
func (s *Shape) PartitionChildren() ([]*Shape, []*Shape) {
	if s.LocalShape.localType() != "Group" {
		// TODO: can we move this to just Group logic instead of all Shape logic?
		// TODO: return error
		panic("Should never AddChildren() to non-Group!\n")
	}

	// HACK: sucks to have to create new Group. When we test out replacing
	// value LocalShape with *LocalShape, we can just alter the Group
	// directly? Otherwise, generalize this into a Copy() method for at least Group.
	g := s.LocalShape.(Group)
	bounds := g.LocalBounds()
	lBounds, rBounds := bounds.SplitBounds()
	l, r := []*Shape{}, []*Shape{}

	newChildren := []*Shape{}
	for _, child := range g.Children {
		if lBounds.ContainsBox(child.ParentSpaceBounds()) {
			l = append(l, child)
		} else if rBounds.ContainsBox(child.ParentSpaceBounds()) {
			r = append(r, child)
		} else {
			newChildren = append(newChildren, child)
		}
		g.Children = newChildren
		s.LocalShape = g
	}

	return l, r
}

// NB: this returns true for "regular shape includes itself"
func (s *Shape) Includes(s2 *Shape) bool {
	if s.IsEqualTo(s2) {
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

func (s *Shape) Bounds() BoundingBox {
	return s.LocalShape.LocalBounds()
}

func (s *Shape) ParentSpaceBounds() BoundingBox {
	b := s.Bounds()
	return b.Transform(s.Transform)
}

func (s *Shape) String() string {
	return fmt.Sprintf(
		"Shape(\n  Label: %v\n  LocalShape: %v\n  Transform: %v\n  Material: %v\n  Parent: %T\n  Shadows: %v\n)",
		s.Label,
		s.LocalShape.localString(),
		s.Transform,
		s.Material,
		s.Parent,
		s.Shadows,
	)
}

func (s *Shape) IsEqualTo(s2 *Shape) bool {
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
