package raytracer

import (
	"math"
)

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

func NewColor(x, y, z float64) Color {
  return Color{x, y, z}
}

func (c *Color) IsEqualTo(c2 Color) bool {
	const tolerance = 0.00001
	equals := func(x, y float64) bool {
		diff := math.Abs(x - y)
    return diff < tolerance
	}

	return equals(c.Red, c2.Red) && equals(c.Green, c2.Green) && equals(c.Blue, c2.Blue)
}

func (c *Color) Add(c2 Color) Color {
  return Color{c.Red + c2.Red, c.Green + c2.Green, c.Blue + c2.Blue}
}

func (c *Color) Subtract(c2 Color) Color {
  return Color{c.Red - c2.Red, c.Green - c2.Green, c.Blue - c2.Blue}
}

func (c *Color) Multiply(scalar float64) Color {
  return Color{c.Red * scalar, c.Green * scalar, c.Blue * scalar}
}

// Returns the Hadamard product (or Schur product) of two colors.
func (c *Color) MultiplyColor(c2 Color) Color {
  return NewColor(c.Red * c2.Red, c.Green * c2.Green, c.Blue * c2.Blue)
}