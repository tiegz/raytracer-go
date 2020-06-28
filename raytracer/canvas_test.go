package raytracer

import (
	"errors"
	"fmt"
	"os"
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
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`

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

func TestCanvasToPpmWithTrailingNewline(t *testing.T) {
	c1 := NewCanvas(5, 3)
	ppm := c1.ToPpm()
	actual := ppm[len(ppm)-1:]
	expected := "\n"

	assertEqualString(t, expected, actual)
}

func TestCanvasSavePpm(t *testing.T) {
	c := NewCanvas(5, 3)
	filepath := "test_file_canvas_save_ppm.ppm"
	err := c.SavePpm(filepath)
	defer os.Remove(filepath)

	assertNil(t, err)
	assertFileExists(t, filepath)
}

func TestCanvasSaveJPEG(t *testing.T) {
	c := NewCanvas(5, 3)
	filepath := "test_file_canvas_save_jpg.jpg"
	err := c.SaveJPEG(filepath)
	defer os.Remove(filepath)

	assertNil(t, err)
	assertFileExists(t, filepath)
}

func TestCanvasSavePNG(t *testing.T) {
	c := NewCanvas(5, 3)
	filepath := "test_file_canvas_save_jpg.png"
	err := c.SavePNG(filepath)
	defer os.Remove(filepath)

	assertNil(t, err)
	assertFileExists(t, filepath)
}

func TestCanvasSaveGIF(t *testing.T) {
	c := NewCanvas(5, 3)
	filepath := "test_file_canvas_save_jpg.gif"
	err := c.SaveGIF(filepath)
	defer os.Remove(filepath)

	assertNil(t, err)
	assertFileExists(t, filepath)
}

func TestReadingAFileWithTheWrongMagicNumber(t *testing.T) {
	ppm := `P32
	1 1
	255
	0 0 0
	`
	_, err := NewCanvasFromPpm(ppm)

	assertEqualError(t, errors.New("raytracer.NewCanvasFromPpm: invalid ppm file, started with P32 instead of P3."), err)
}

func TestReadingAPpmReturnsACanvasOfTheRightSize(t *testing.T) {
	ppm := `P3
	10 2
	255
	0 0 0  0 0 0  0 0 0  0 0 0  0 0 0
	0 0 0  0 0 0  0 0 0  0 0 0  0 0 0
	0 0 0  0 0 0  0 0 0  0 0 0  0 0 0
	0 0 0  0 0 0  0 0 0  0 0 0  0 0 0
	`
	c, err := NewCanvasFromPpm(ppm)
	assertNil(t, err)
	assertEqualInt(t, 10, c.Width)
	assertEqualInt(t, 2, c.Height)
}

func TestReadingPixelDataFromAPpmFile(t *testing.T) {
	ppm := `P3
	4 3
	255
	255 127 0  0 127 255  127 255 0  255 255 255
	0 0 0  255 0 0  0 255 0  0 0 255
	255 255 0  0 255 255  255 0 255  127 127 127
	`
	c, _ := NewCanvasFromPpm(ppm)
	testCases := []struct {
		x     int
		y     int
		color Color
	}{
		{0, 0, NewColor(1, 0.49803, 0)},
		{1, 0, NewColor(0, 0.49803, 1)},
		{2, 0, NewColor(0.49803, 1, 0)},
		{3, 0, NewColor(1, 1, 1)},
		{0, 1, NewColor(0, 0, 0)},
		{1, 1, NewColor(1, 0, 0)},
		{2, 1, NewColor(0, 1, 0)},
		{3, 1, NewColor(0, 0, 1)},
		{0, 2, NewColor(1, 1, 0)},
		{1, 2, NewColor(0, 1, 1)},
		{2, 2, NewColor(1, 0, 1)},
		{3, 2, NewColor(0.49803, 0.49803, 0.49803)},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("testCases[%d]", idx), func(t *testing.T) {
			assertEqualColor(t, tc.color, c.PixelAt(tc.x, tc.y))
		})
	}
}

func TestPpmParsingIgnoresCommentsLines(t *testing.T) {
	ppm := `P3
	# this is a comment
	2 1
	# this, too
	255
	# another comment
	255 255 255
	# oh, no, comments in the pixel data!
	255 0 255
	`
	c, _ := NewCanvasFromPpm(ppm)
	assertEqualColor(t, NewColor(1, 1, 1), c.PixelAt(0, 0))
	assertEqualColor(t, NewColor(1, 0, 1), c.PixelAt(1, 0))
}

func TestPpmParsingAllowsAnRgbTripleToSpanLines(t *testing.T) {
	ppm := `P3
    1 1
    255
    51
    153

    204
	`
	c, _ := NewCanvasFromPpm(ppm)
	assertEqualColor(t, NewColor(0.2, 0.6, 0.8), c.PixelAt(0, 0))
}

func TestPpmParsingRespectsTheScaleSetting(t *testing.T) {
	ppm := `P3
    2 2
    100
    100 100 100  50 50 50
    75 50 25  0 0 0
	`
	c, _ := NewCanvasFromPpm(ppm)
	assertEqualColor(t, NewColor(0.75, 0.5, 0.25), c.PixelAt(0, 1))
}

/////////////
// Benchmarks
/////////////

func BenchmarkCanvasMethodIsEqualTo(b *testing.B) {
	c1 := NewCanvas(100, 100)
	for i := 0; i < b.N; i++ {
		c1.IsEqualTo(c1)
	}
}

func BenchmarkCanvasMethodToPpm(b *testing.B) {
	c1 := NewCanvas(10, 10)
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			c1.WritePixel(x, y, Color{1, 1, 1})
		}
	}
	for i := 0; i < b.N; i++ {
		c1.ToPpm()
	}
}
