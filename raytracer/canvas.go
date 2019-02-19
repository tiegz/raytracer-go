package raytracer

import (
	"math"
)

type Canvas struct {
	Width  float64
	Height float64
}

func NewCanvas(w, h float64) Canvas {
  return Canvas{w, h}
}

func (c *Canvas) IsEqualTo(c2 Canvas) bool {
	const tolerance = 0.00001
	equals := func(x, y float64) bool {
		diff := math.Abs(x - y)
    return diff < tolerance
	}

	return equals(c.Width, c2.Width) && equals(c.Height, c2.Height)
}
