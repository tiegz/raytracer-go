package raytracer

type Canvas struct {
	Width  int
	Height int
	Pixels []Color
}

func NewCanvas(w, h int) Canvas {
	return Canvas{w, h, make([]Color, h*w)}
}

func (c *Canvas) IsEqualTo(c2 Canvas) bool {
  if (c.Width != c2.Width) {
    return false
  }
  if (c.Height != c2.Height) {
    return false
  }

  for x := 0; x < c.Width; x = x + 1 {
    for y := 0; y < c.Height; y = y + 1 {
      if c.PixelAt(x, y) != c2.PixelAt(x, y){
        return false
      }
    }
  }

	return true
}

func (c *Canvas) WritePixel(x, y int, color Color) {
	index := c.Height*y + x
	c.Pixels[index] = color
}

func (c *Canvas) PixelAt(x, y int) Color {
	index := c.Height*y + x
	return c.Pixels[index]
}
