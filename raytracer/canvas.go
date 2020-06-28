package raytracer

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Canvas struct {
	Width      int
	Height     int
	ColorScale float64
	Pixels     []Color
}

func NewCanvas(w, h int, defaultColor ...Color) Canvas {
	c := Canvas{w, h, 255, make([]Color, h*w)}

	if len(defaultColor) > 0 {
		for x := 0; x < c.Width; x += 1 {
			for y := 0; y < c.Height; y += 1 {
				c.WritePixel(x, y, defaultColor[0])
			}
		}
	}

	return c
}

func NewCanvasFromPpm(ppm string) (Canvas, error) {
	var colorScale float64
	var w, h int
	var err error
	var c Canvas
	buf := bytes.NewBuffer([]byte(ppm))
	scanner := bufio.NewScanner(bufio.NewReader(buf))
	var lines []string

	// Read lines first
	for scanner.Scan() {
		currentLine := strings.TrimSpace(string(scanner.Text()))
		if !strings.HasPrefix(currentLine, "#") {
			lines = append(lines, currentLine)
		}
	}

	// Check magic number
	magicNumber, lines := lines[0], lines[1:]
	if magicNumber != "P3" {
		return c, fmt.Errorf("raytracer.NewCanvasFromPpm: invalid ppm file, started with %s instead of P3.", magicNumber)
	}

	// Set dimensions
	dimensions, lines := lines[0], lines[1:]
	if _, err := fmt.Sscanf(dimensions, "%d %d", &w, &h); err != nil {
		return c, err
	}

	// Set scale
	scale, lines := lines[0], lines[1:]
	if _, err := fmt.Sscanf(scale, "%f", &colorScale); err != nil {
		return c, err
	}

	c = Canvas{w, h, colorScale, make([]Color, h*w)}
	// Pixel triplets aren't necessarily grouped into rows by line, so we just have to take in a single-dimensional array instead and read that.
	pixelData := strings.Fields(strings.Join(lines, " "))
	pixelCount := c.Width * c.Height
	var r, g, b int

	for i := 0; i < pixelCount; i++ {
		if r, err = strconv.Atoi(string(pixelData[(i * 3)])); err != nil {
			return c, err
		}
		if g, err = strconv.Atoi(string(pixelData[(i*3)+1])); err != nil {
			return c, err
		}
		if b, err = strconv.Atoi(string(pixelData[(i*3)+2])); err != nil {
			return c, err
		}
		x, y := i%c.Width, i/c.Width
		c.WritePixel(x, y, NewColor(float64(r)/c.ColorScale, float64(g)/c.ColorScale, float64(b)/c.ColorScale))
	}

	return c, nil
}

func (c Canvas) String() string {
	return fmt.Sprintf("Canvas(\nWidth: %v\nHeight: %v\nColorScale: %v\n)", c.Width, c.Height, c.ColorScale)
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
	if x >= 0 && x < c.Width && y >= 0 && y < c.Height {
		index := (c.Width * y) + x
		c.Pixels[index] = color
	} else {
		fmt.Printf("Warning: skipping WritePixel(%d, %d) because it's outside of canvas (%d, %d)\n", x, y, c.Width, c.Height)
	}
}

func (c *Canvas) PixelAt(x, y int) Color {
	index := (c.Width * y) + x
	return c.Pixels[index]
}

func (c *Canvas) SaveJPEG(filepath string) error {
	target := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
	for x := 0; x < c.Width; x++ {
		for y := 0; y < c.Height; y++ {
			target.Set(x, y, c.PixelAt(x, y))
		}
	}

	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	err = jpeg.Encode(f, target, nil)
	return err
}

func (c *Canvas) SavePpm(filepath string) error {
	ppm := c.ToPpm()
	ppmBytes := []byte(ppm)
	err := ioutil.WriteFile(filepath, ppmBytes, 0644)
	return err
}

func (c *Canvas) ToPpm() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("P3\n%d %d\n%d\n", c.Width, c.Height, int(c.ColorScale)))

	for i, v := range c.Pixels {
		r, g, b := v.ScaledRGB(c.ColorScale)
		buffer.WriteString(fmt.Sprintf("%d %d %d", r, g, b))

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
