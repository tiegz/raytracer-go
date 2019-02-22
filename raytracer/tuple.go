package raytracer

import (
	"fmt"
	"math"
)

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64 // 1: point, 0: vector
}

func (t *Tuple) Type() string {
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
		return fmt.Sprintf("Vector( %.3f, %.3f, %.3f, %.3f )", t.X, t.Y, t.Z, t.W)
	} else {
		return fmt.Sprintf("Point( %.3f, %.3f, %.3f, %.3f )", t.X, t.Y, t.Z, t.W)
	}
}

func (t *Tuple) IsEqualTo(t2 Tuple) bool {
	const tolerance = 0.00001
	equals := func(x, y float64) bool {
		diff := math.Abs(x - y)
		return diff < tolerance
	}

	return equals(t.X, t2.X) && equals(t.Y, t2.Y) && equals(t.Z, t2.Z) && equals(t.W, t2.W)
}

func (t *Tuple) Add(t2 Tuple) Tuple {
	return Tuple{t.X + t2.X, t.Y + t2.Y, t.Z + t2.Z, t.W + t2.W}
}

func (t *Tuple) Subtract(t2 Tuple) Tuple {
	return Tuple{t.X - t2.X, t.Y - t2.Y, t.Z - t2.Z, t.W - t2.W}
}

func (t *Tuple) Negate() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

func (t *Tuple) Multiply(scalar float64) Tuple {
	return Tuple{t.X * scalar, t.Y * scalar, t.Z * scalar, t.W * scalar}
}

func (t *Tuple) Divide(scalar float64) Tuple {
	return Tuple{t.X / scalar, t.Y / scalar, t.Z / scalar, t.W / scalar}
}

// Returns the magnitude of a vector.
func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2) + math.Pow(t.W, 2))
}

// Returns a normalized vector.
func (t *Tuple) Normalized() Tuple {
	mag := t.Magnitude()

	return Tuple{t.X / mag, t.Y / mag, t.Z / mag, t.W / mag}
}

// Returns the dot product (aka scalar product or inner product) of two vectors.
// "... the smaller the dot product, the larger the angle between them ..."
func (t *Tuple) Dot(t2 Tuple) float64 {
	return (t.X * t2.X) + (t.Y * t2.Y) + (t.Z * t2.Z) + (t.W * t2.W)
}

// Returns the cross product of two vectors.
// "... gives you a new vector that is perpendicular to both of the original vectors ..."
func (t *Tuple) Cross(t2 Tuple) Tuple {
	return NewVector(
		(t.Y*t2.Z)-(t.Z*t2.Y),
		(t.Z*t2.X)-(t.X*t2.Z),
		(t.X*t2.Y)-(t.Y*t2.X),
	)
}
