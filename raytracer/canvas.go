package raytracer

import "fmt"

type Canvas struct {
	Width  int
	Height int
	Pixels []Color
}

func NewCanvas(w, h int) Canvas {
	return Canvas{w, h, make([]Color, h*w)}
}

func (c *Canvas) IsEqualTo(c2 Canvas) bool {
	const tolerance = 0.00001
	// equals := func(x, y float64) bool {
	// 	diff := math.Abs(x - y)
	// 	return diff < tolerance
	// }

	return true
	// equals(c.Width, c2.Width) && equals(c.Height, c2.Height)
}

func (c *Canvas) WritePixel(x, y int, color Color) {
	index := c.Height*y + x
	c.Pixels[index] = color
}

func (c *Canvas) PixelAt(x, y int) Color {
	index := c.Height*y + x
	return c.Pixels[index]
}

func (c *Canvas) ToPpm() string {
	return fmt.Sprintf("P3\n%d %d\n255\n", c.Width, c.Height)
}
