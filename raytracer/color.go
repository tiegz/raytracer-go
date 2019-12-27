package raytracer

import (
	"fmt"
)

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

func NewColor(x, y, z float64) Color {
	return Color{x, y, z}
}

var Colors = map[string]Color{
	"White":      NewColor(1, 1, 1),
	"Gray":       NewColor(0.5, 0.5, 0.5),
	"Black":      NewColor(0, 0, 0),
	"Red":        NewColor(1, 0, 0),
	"Orange":     NewColor(1, 0.5, 0),
	"Yellow":     NewColor(1, 1, 0),
	"Green":      NewColor(0, 1, 0),
	"Blue":       NewColor(0, 0, 1),
	"Purple":     NewColor(0.5, 0, 0.5),
	"DarkGray":   NewColor(0.5, 0.5, 0.5),
	"DarkBlack":  NewColor(0, 0, 0),
	"DarkRed":    NewColor(1, 0, 0),
	"DarkOrange": NewColor(1, 0.5, 0),
	"DarkYellow": NewColor(1, 1, 0),
	"DarkGreen":  NewColor(0, 1, 0),
	"DarkBlue":   NewColor(0, 0, 1),
	"DarkPurple": NewColor(0.5, 0, 0.5),
}

func (c Color) IsEqualTo(c2 Color) bool {
	return equalFloat64s(c.Red, c2.Red) && equalFloat64s(c.Green, c2.Green) && equalFloat64s(c.Blue, c2.Blue)
}

func (c Color) String() string {
	return fmt.Sprintf("Color( %v %v %v )", c.Red, c.Green, c.Blue)
}

func (c Color) Add(c2 Color) Color {
	return Color{c.Red + c2.Red, c.Green + c2.Green, c.Blue + c2.Blue}
}

func (c Color) Subtract(c2 Color) Color {
	return Color{c.Red - c2.Red, c.Green - c2.Green, c.Blue - c2.Blue}
}

func (c Color) Multiply(scalar float64) Color {
	return Color{c.Red * scalar, c.Green * scalar, c.Blue * scalar}
}

func (c Color) Divide(scalar float64) Color {
	return Color{c.Red / scalar, c.Green / scalar, c.Blue / scalar}
}

// Returns the Hadamard product (or Schur product) of two colors.
func (c Color) MultiplyColor(c2 Color) Color {
	return NewColor(c.Red*c2.Red, c.Green*c2.Green, c.Blue*c2.Blue)
}
