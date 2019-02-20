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

func TestCanvasToPpmWithDataExceeding70Characters(t *testing.T) {
	c1 := NewCanvas(10, 2, Color{1, 0.8, 0.6})
	lines := strings.Split(c1.ToPpm(), "\n")
	actual := strings.Join(lines[3:7], "\n")
	expected := `255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153`

	assertEqualString(t, expected, actual)
}

// Scenario: Splitting long lines in PPM files
// Given c ← canvas(10, 2)
// When every pixel of c is set to color(1, 0.8, 0.6)
// And ppm ← canvas_to_ppm(c) Prepared exclusively for Tieg Zaharia
// report erratum • discuss
// Saving a Canvas • 21

// Then lines 4-7 of ppm are """
// """
