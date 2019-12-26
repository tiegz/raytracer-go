package raytracer

import (
	"fmt"
	"math"
)

// TODO could Tuple be an interface and Point/Vector types that implement it?
type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64 // 1: point, 0: vector
}

func (t Tuple) Type() string {
	if t.W == 1.0 {
		return "Point"
	} else {
		return "Vector"
	}
}

func NewPoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func NewVector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func (t Tuple) String() string {
	if t.W == 0 {
		return fmt.Sprintf("Vector( %.5f, %.5f, %.5f )", t.X, t.Y, t.Z)
	} else {
		return fmt.Sprintf("Point( %.5f, %.5f, %.5f )", t.X, t.Y, t.Z)
	}
}

func (t Tuple) IsEqualTo(t2 Tuple) bool {
	return equalFloat64s(t.X, t2.X) && equalFloat64s(t.Y, t2.Y) && equalFloat64s(t.Z, t2.Z) && equalFloat64s(t.W, t2.W)
}

func (t Tuple) Add(t2 Tuple) Tuple {
	return Tuple{t.X + t2.X, t.Y + t2.Y, t.Z + t2.Z, t.W + t2.W}
}

func (t Tuple) Subtract(t2 Tuple) Tuple {
	return Tuple{t.X - t2.X, t.Y - t2.Y, t.Z - t2.Z, t.W - t2.W}
}

func (t Tuple) Negate() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

func (t Tuple) Multiply(scalar float64) Tuple {
	return Tuple{t.X * scalar, t.Y * scalar, t.Z * scalar, t.W * scalar}
}

func (t Tuple) Divide(scalar float64) Tuple {
	return Tuple{t.X / scalar, t.Y / scalar, t.Z / scalar, t.W / scalar}
}

// Return the reflection of this vector, off a given normal.
func (t Tuple) Reflect(normal Tuple) Tuple {
	reflection := normal.Multiply(2)
	reflection = reflection.Multiply(t.Dot(normal))
	reflection = t.Subtract(reflection)
	return reflection
}

// Returns the magnitude of a vector.
func (t Tuple) Magnitude() float64 {
	return math.Sqrt((t.X * t.X) + (t.Y * t.Y) + (t.Z * t.Z) + (t.W * t.W))
}

// Returns a normalized vector.
func (t Tuple) Normalized() Tuple {
	mag := t.Magnitude()

	return Tuple{t.X / mag, t.Y / mag, t.Z / mag, t.W / mag}
}

// Returns the dot product (aka scalar product or inner product) of two vectors.
// "... the smaller the dot product, the larger the angle between them ..."
func (t Tuple) Dot(t2 Tuple) float64 {
	return (t.X * t2.X) + (t.Y * t2.Y) + (t.Z * t2.Z) + (t.W * t2.W)
}

// Returns the cross product of two vectors.
// "... gives you a new vector that is perpendicular to both of the original vectors ..."
func (t Tuple) Cross(t2 Tuple) Tuple {
	return NewVector(
		(t.Y*t2.Z)-(t.Z*t2.Y),
		(t.Z*t2.X)-(t.X*t2.Z),
		(t.X*t2.Y)-(t.Y*t2.X),
	)
}
