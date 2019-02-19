package raytracer

import (
	"testing"
)

func TestNewCanvas(t *testing.T) {
	c1 := NewCanvas(10, 20)

	assertEqualInt(t, 10, c1.Width)
	assertEqualInt(t, 20, c1.Height)
}

func TestWritePixel(t *testing.T) {
	c1 := NewCanvas(10, 20)
	red := NewColor(1, 0, 0)

	c1.WritePixel(2, 3, red)
	assertEqualColor(t, red, c1.PixelAt(2, 3))
}
