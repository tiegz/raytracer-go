package raytracer

import (
	"bytes"
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
				c.WritePixel(x, y, defaultColor[0])
			}
		}
	}

	return c
}

func (c *Canvas) IsEqualTo(c2 Canvas) bool {
	if c.Width != c2.Width || c.Height != c2.Height {
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
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", c.Width, c.Height))

	for i, v := range c.Pixels {
		buffer.WriteString(
			fmt.Sprintf("%d %d %d",
				int(math.Ceil(math.Min(255, math.Max(0, v.Red*255)))),
				int(math.Ceil(math.Min(255, math.Max(0, v.Green*255)))),
				int(math.Ceil(math.Min(255, math.Max(0, v.Blue*255)))),
			))

		if (i != 0) && (i+1)%(c.Width) == 0 {
			// NB this still happens on the last line, to ensure there's
			// a trailing newline in the file.
			buffer.WriteString("\n")
		} else {
			buffer.WriteString(" ")
		}
	}

	data := buffer.String()

	// " ... no line in a PPM file should be more than 70 characters long ..."
	lines := strings.Split(data, "\n")
	bufferTruncated := bytes.Buffer{}

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		newLineBuffer := bytes.Buffer{}
		newLineBufferLength := 0

		for _, token := range tokens {
			if newLineBufferLength+len(token)+1 > 70 {
				bufferTruncated.WriteString(newLineBuffer.String())
				bufferTruncated.WriteString("\n")
				newLineBuffer = bytes.Buffer{}
				newLineBuffer.WriteString(token)
				newLineBufferLength = len(token)
			} else {
				if newLineBufferLength > 0 {
					newLineBuffer.WriteString(" ")
					newLineBufferLength += 1
				}
				newLineBuffer.WriteString(token)
				newLineBufferLength += len(token)
			}
		}
		bufferTruncated.WriteString(newLineBuffer.String())
		bufferTruncated.WriteString("\n")
	}

	return bufferTruncated.String()
}
