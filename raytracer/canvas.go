package raytracer

import (
	"fmt"
	"math"
	"strings"
)

type Canvas struct {
	Width  int
	Height int
	Pixels []Color
}

func NewCanvas(w, h int, defaultColor ...Color) Canvas {
	c := Canvas{w, h, make([]Color, h*w)}

	if len(defaultColor) > 0 {
		for x := 0; x < c.Width; x += 1 {
			for y := 0; y < c.Height; y += 1 {
				c.WritePixel(x, y, Color{1, 0.8, 0.6})
			}
		}
	}

	return c
}

func (c *Canvas) IsEqualTo(c2 Canvas) bool {
	if c.Width != c2.Width {
		return false
	}
	if c.Height != c2.Height {
		return false
	}

	for x := 0; x < c.Width; x = x + 1 {
		for y := 0; y < c.Height; y = y + 1 {
			if c.PixelAt(x, y) != c2.PixelAt(x, y) {
				return false
			}
		}
	}

	return true
}

func (c *Canvas) WritePixel(x, y int, color Color) {
	index := (c.Width * y) + x
	c.Pixels[index] = color
}

func (c *Canvas) PixelAt(x, y int) Color {
	index := (c.Width * y) + x
	return c.Pixels[index]
}

func (c *Canvas) ToPpm() string {
	data := fmt.Sprintf("P3\n%d %d\n255\n", c.Width, c.Height)
	currentLineLength := 0
	for i, v := range c.Pixels {
		red := fmt.Sprintf("%d", int(math.Ceil(math.Min(255, math.Max(0, v.Red*255)))))
		green := fmt.Sprintf("%d", int(math.Ceil(math.Min(255, math.Max(0, v.Green*255)))))
		blue := fmt.Sprintf("%d", int(math.Ceil(math.Min(255, math.Max(0, v.Blue*255)))))

		addToken := func(token, data string, currentLineLength int) (string, int) {
			// if (currentLineLength + len(token) > 70) {
			// 	data += "\n"
			// 	currentLineLength = 0
			// }

			data += token
			currentLineLength += len(token)

			return data, currentLineLength
		}

		data, currentLineLength = addToken(red, data, currentLineLength)
		data, currentLineLength = addToken(" ", data, currentLineLength)
		data, currentLineLength = addToken(green, data, currentLineLength)
		data, currentLineLength = addToken(" ", data, currentLineLength)
		data, currentLineLength = addToken(blue, data, currentLineLength)
		if (i != 0) && (i+1)%(c.Width) == 0 {
			// NB this still happens on the last line, to ensure there's
			// a trailing newline in the file.
			data += "\n"
			currentLineLength = 0
		} else {
			data, currentLineLength = addToken(" ", data, currentLineLength)
		}
	}

	// " ... no line in a PPM file should be more than 70 characters long ..."
	lines := strings.Split(data, "\n")
	dataTruncatedTo70Chars := ""
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		newLine := ""
		for _, token := range tokens {
			if len(newLine)+len(token)+1 > 70 {
				dataTruncatedTo70Chars += newLine + "\n"
				newLine = token
			} else {
				if len(newLine) > 0 {
					newLine += " "
				}
				newLine += token
			}
		}
		dataTruncatedTo70Chars += newLine + "\n"
	}

	return dataTruncatedTo70Chars
}
