package raytracer

import (
	"strings"
	"testing"
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
	actual := strings.Join(strings.Split(c1.ToPpm(), "\n")[0:3], "\n")
	expected := "P3\n5 3\n255"

	assertEqualString(t, expected, actual)
}

func TestCanvasToPpmWithData(t *testing.T) {
	c1 := NewCanvas(5, 3)
	c1.WritePixel(0, 0, Color{1.5, 0, 0})
	c1.WritePixel(2, 1, Color{0, 0.5, 0})
	c1.WritePixel(4, 2, Color{-0.5, 0, 1})
	lines := strings.Split(c1.ToPpm(), "\n")

	actual := strings.Join(lines[3:len(lines)-1], "\n")
	expected := `255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255`
	assertEqualString(t, expected, actual)
}

// Scenario: Constructing the PPM pixel data
// Given c ← canvas(5, 3)
// And c1 ← color(1.5, 0, 0)
// And c2 ← color(0, 0.5, 0)
// And c3 ← color(-0.5, 0, 1)
// When write_pixel(c, 0, 0, c1)
// And write_pixel(c, 2, 1, c2)
// And write_pixel(c, 4, 2, c3)
// And ppm ← canvas_to_ppm(c)
// Then lines 4-6 of ppm are """
//                      """
