package raytracer

import (
  "math"
)

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64 // 0: point, 1: vector
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

func (t *Tuple) IsEqualTo(t2 Tuple) bool {
  const tolerance = 0.00001
  equals := func(x, y float64) bool {
    diff := math.Abs(x - y)
    return diff < tolerance
  }

  return equals(t.X, t2.X) && equals(t.Y, t2.Y) && equals(t.Z, t2.Z) && equals(t.W, t2.W)
}

func (t *Tuple) Add(t2 Tuple) Tuple {
  return t2
}
