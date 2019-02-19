package raytracer

import (
	"testing"
)

func TestNewCanvas(t *testing.T) {
  c1 := NewCanvas(10, 20)

  assertEqualFloat64(t, 10, c1.Width)
  assertEqualFloat64(t, 20, c1.Height)
}

func TestWritePixel(t *testing.T) {
  c1 := NewCanvas(10, 20)
  red := NewColor(1, 0, 0)

  c1.WritePixel(c1, 2, 3, red)
  assertEqualColor(red, c1.PixelAt(2, 3))
}
