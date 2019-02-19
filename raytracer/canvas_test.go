package raytracer

import (
	"testing"
  "strings"
)

func TestNewCanvas(t *testing.T) {
  c1 := NewCanvas(10, 20)
  expectedColor := NewColor(0, 0, 0)

  assertEqualInt(t, 10, c1.Width)
  assertEqualInt(t, 20, c1.Height)

  for x := 0; x < 10; x = x + 1 {
    for y := 20; y < 20; y = y + 1 {
      assertEqualColor(t, expectedColor, c1.PixelAt(x, y))
    }
  }
}

func TestWritePixel(t *testing.T) {
	c1 := NewCanvas(10, 20)
	red := NewColor(1, 0, 0)

	c1.WritePixel(2, 3, red)
	assertEqualColor(t, red, c1.PixelAt(2, 3))
}

func TestCanvasToPpm(t *testing.T) {
  c1 := NewCanvas(5, 3)
  actual := strings.Join(strings.Split(c1.ToPpm(), "\n")[0:2], "\n")
  expected := "P3\n5\n3"

  assertEqualString(t, expected, actual)
}
