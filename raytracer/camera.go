package raytracer

import (
	"fmt"
	"math"
)

type Camera struct {
	hSize            int
	vSize            int
	halfWidth        float64
	halfHeight       float64
	fieldOfView      float64
	pixelSize        float64
	transform        Matrix
	inverseTransform Matrix
}

// NewCamera returns a Camera, which renders a canvas 1 unit in front of it.
//
// The h argument is the horizontal size of the canvas.
// The v argument is the vertical size of the canvas.
// The f argument is the field-of-view of the camera. (i.e. smaller means zoomed-in)
//
// The Camera also has a Transform attribute, describing the world's orientation relative to the camera.
func NewCamera(h, v int, f float64) Camera {
	c := Camera{}
	c.SetSize(h, v, f)
	c.SetTransform(IdentityMatrix())
	return c
}

func (c Camera) String() string {
	return fmt.Sprintf("Camera(\n  HSize: %d\n  Vsize: %d\n Field of View: %f\n  Transform: %v\n)", c.GetHSize(), c.GetVSize(), c.GetFieldOfView(), c.GetTransform())
}

func (c *Camera) GetHSize() int               { return c.hSize }
func (c *Camera) GetVSize() int               { return c.vSize }
func (c *Camera) GetFieldOfView() float64     { return c.fieldOfView }
func (c *Camera) GetPixelSize() float64       { return c.pixelSize }
func (c *Camera) GetTransform() Matrix        { return c.transform }
func (c *Camera) GetInverseTransform() Matrix { return c.inverseTransform }

func (c *Camera) SetSize(h, v int, f float64) {
	c.hSize = h
	c.vSize = v
	c.fieldOfView = f

	// Calculate pixel size
	halfView := math.Tan(c.fieldOfView / 2) // p 106 illustration
	aspectRatio := float64(c.GetHSize()) / float64(c.GetVSize())
	if aspectRatio >= 1 { // h >= v
		c.halfWidth = halfView
		c.halfHeight = halfView / aspectRatio
	} else { // v > h
		c.halfWidth = halfView * aspectRatio
		c.halfHeight = halfView
	}
	c.pixelSize = (c.halfWidth * 2) / float64(c.GetHSize())
}

// TODO: make transform private, to force usage of SetTransform() instead?
func (c *Camera) SetTransform(m Matrix) {
	c.transform = m
	c.inverseTransform = m.Inverse()
}

// RayForPixel returns a ray, from the camera through the point indicated.
func (c *Camera) RayForPixel(pixelX, pixelY int) Ray {
	// ... the offset from the edge of the canvas to the pixel's center ...
	xOffset := (float64(pixelX) + 0.5) * c.pixelSize
	yOffset := (float64(pixelY) + 0.5) * c.pixelSize

	// ... the untransformed coordinates of the pixel in world-space. ...
	// ... (remember that the camera looks toward -z, so +x is to the *left*.) ...
	worldX := c.halfWidth - xOffset
	worldY := c.halfHeight - yOffset

	// ... transform the canvas point and the origin, and then compute the ray's direction vector. ...
	// ... (remember that the canvas is at z=-1) ...
	pixel := c.inverseTransform.MultiplyByTuple(NewPoint(worldX, worldY, -1))
	origin := c.inverseTransform.MultiplyByTuple(NewPoint(0, 0, 0))

	direction := pixel.Subtract(origin)
	direction = direction.Normalized()

	return NewRay(origin, direction)
}

// Returns a canvas that renders the world from the given camera.
func (c *Camera) Render(w World) Canvas {
	canvas := NewCanvas(c.hSize, c.vSize)

	for y := 0; y < c.vSize; y += 1 {
		for x := 0; x < c.hSize; x += 1 {
			r := c.RayForPixel(x, y)
			color := w.ColorAt(r, DefaultMaximumReflections)
			canvas.WritePixel(x, y, color)
		}
	}

	return canvas
}

// Same as Render(), while also outputting the current number of pixels rendered to stdout.
func (c *Camera) RenderWithProgress(w World) Canvas {
	canvas := NewCanvas(c.hSize, c.vSize)

	count := c.hSize * c.vSize
	i := 0

	for y := 0; y < c.vSize; y += 1 {
		for x := 0; x < c.hSize; x += 1 {
			r := c.RayForPixel(x, y)
			color := w.ColorAt(r, DefaultMaximumReflections)
			canvas.WritePixel(x, y, color)

			i += 1
			progress := ((float64(i) / float64(count)) * 100)
			fmt.Printf("\rProgress: %6.02f%%", progress)
		}
	}
	fmt.Println()

	return canvas
}
