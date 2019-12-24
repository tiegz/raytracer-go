package raytracer

import (
	"fmt"
	"math"
)

type Camera struct {
	HSize       int
	VSize       int
	HalfWidth   float64
	HalfHeight  float64
	FieldOfView float64
	Transform   Matrix
}

// NewCamera returns a Camera, which renders a canvas 1 unit in front of it.
//
// The h argument is the horizontal size of the canvas.
// The v argument is the vertical size of the canvas.
// The f argument is the field-of-view of the camera. (i.e. smaller means zoomed-in)
//
// The Camera also has a Transform attribute, describing the world's orientation relative to the camera.
func NewCamera(h, v int, f float64) Camera {
	return Camera{
		HSize:       h,
		VSize:       v,
		FieldOfView: f,
		Transform:   IdentityMatrix(),
	}
}

func (c *Camera) PixelSize() float64 {
	halfView := math.Tan(c.FieldOfView / 2) // p 106 illustration
	aspectRatio := float64(c.HSize) / float64(c.VSize)

	// TODO can we move these value settings into the constructor?
	if aspectRatio >= 1 { // h >= v
		c.HalfWidth = halfView
		c.HalfHeight = halfView / aspectRatio
	} else { // v > h
		c.HalfWidth = halfView * aspectRatio
		c.HalfHeight = c.HalfHeight
	}

	return (c.HalfWidth * 2) / float64(c.HSize)
}

// TODO memoize PixelSize() for this func?
// RayForPixel returns a ray, from the camera through the point indicated.
func (c *Camera) RayForPixel(pixelX, pixelY int) Ray {
	// ... the offset from the edge of the canvas to the pixel's center ...
	xOffset := (float64(pixelX) + 0.5) * c.PixelSize()
	yOffset := (float64(pixelY) + 0.5) * c.PixelSize()

	// ... the untransformed coordinates of the pixel in world-space. ...
	// ... (remember that the camera looks toward -z, so +x is to the *left*.) ...
	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset

	// ... transform the canvas point and the origin, and then compute the ray's direction vector. ...
	// ... (remember that the canvas is at z=-1) ...
	inverseCameraTransform := c.Transform.Inverse()
	pixel := inverseCameraTransform.MultiplyByTuple(NewPoint(worldX, worldY, -1))
	origin := inverseCameraTransform.MultiplyByTuple(NewPoint(0, 0, 0))

	direction := pixel.Subtract(origin)
	direction = direction.Normalized()

	return NewRay(origin, direction)
}

// Returns a canvas that renders the world from the given camera.
func (c *Camera) Render(w World) Canvas {
	canvas := NewCanvas(c.HSize, c.VSize)

	for y := 0; y < c.VSize; y += 1 {
		for x := 0; x < c.HSize; x += 1 {
			r := c.RayForPixel(x, y)
			color := w.ColorAt(r, DefaultMaximumReflections)
			canvas.WritePixel(x, y, color)
		}
	}

	return canvas
}

// Same as Render(), while also outputting the current number of pixels rendered to stdout.
func (c *Camera) RenderWithProgress(w World) Canvas {
	canvas := NewCanvas(c.HSize, c.VSize)

	count := c.HSize * c.VSize
	i := 0

	for y := 0; y < c.VSize; y += 1 {
		for x := 0; x < c.HSize; x += 1 {
			r := c.RayForPixel(x, y)
			color := w.ColorAt(r, DefaultMaximumReflections)
			canvas.WritePixel(x, y, color)

			i += 1
			progress := ((float64(i) / float64(count)) * 100)
			fmt.Printf("\rProgress: %6.02f%%", progress)
		}
	}

	return canvas
}
