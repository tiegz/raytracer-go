package raytracer

import (
	"math"
	"testing"
)

func TestConstructingACamera(t *testing.T) {
	hsize := 160
	vsize := 120
	fieldOfView := math.Pi / 2
	camera := NewCamera(hsize, vsize, fieldOfView)

	assertEqualInt(t, 160, camera.HSize)
	assertEqualInt(t, 120, camera.VSize)
	assertEqualFloat64(t, math.Pi/2, camera.FieldOfView)
	assertEqualMatrix(t, IdentityMatrix(), camera.Transform)
}

func TestPixelSizeForHorizontalCanvas(t *testing.T) {
	c := NewCamera(200, 125, math.Pi/2)
	assertEqualFloat64(t, 0.01, c.PixelSize)
}

func TestPixelSizeForVerticalCanvas(t *testing.T) {
	c := NewCamera(125, 200, math.Pi/2)
	assertEqualFloat64(t, 0.01, c.PixelSize)
}

func TestConstructingARayThroughCenterOfTheCanvas(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	r := c.RayForPixel(100, 50)

	assertEqualTuple(t, NewPoint(0, 0, 0), r.Origin)
	assertEqualTuple(t, NewVector(0, 0, -1), r.Direction)
}

func TestConstructingARayThroughACornerOfTheCanvas(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	r := c.RayForPixel(0, 0)

	assertEqualTuple(t, NewPoint(0, 0, 0), r.Origin)
	assertEqualTuple(t, NewVector(0.66519, 0.33259, -0.66851), r.Direction)
}

func TestConstructingARayWhenCameraIsTransformed(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	transform := NewRotateY(math.Pi / 4)
	transform = transform.Multiply(NewTranslation(0, -2, 5))
	c.SetTransform(transform)
	r := c.RayForPixel(100, 50)

	assertEqualTuple(t, NewPoint(0, 2, -5), r.Origin)
	assertEqualTuple(t, NewVector(math.Sqrt(2)/2, 0, -(math.Sqrt(2)/2)), r.Direction)
}

func TestRenderingWorldWithCamera(t *testing.T) {
	w := DefaultWorld()
	c := NewCamera(11, 11, math.Pi/2)
	from := NewPoint(0, 0, -5)
	to := NewPoint(0, 0, 0)
	up := NewVector(0, 1, 0)
	c.SetTransform(NewViewTransform(from, to, up))
	image := c.Render(w, 1, false)

	expected := NewColor(0.38066, 0.47583, 0.2855)
	actual := image.PixelAt(5, 5)
	assertEqualColor(t, expected, actual)
}

func TestRenderingWorldWithCameraInParallel(t *testing.T) {
	w := DefaultWorld()
	c := NewCamera(11, 11, math.Pi/2)
	from := NewPoint(0, 0, -5)
	to := NewPoint(0, 0, 0)
	up := NewVector(0, 1, 0)
	c.SetTransform(NewViewTransform(from, to, up))
	image := c.Render(w, 2, false)

	expected := NewColor(0.38066, 0.47583, 0.2855)
	actual := image.PixelAt(5, 5)
	assertEqualColor(t, expected, actual)
}

/////////////
// Benchmarks
/////////////

func BenchmarkCameraMethodRayForPixel(b *testing.B) {
	// Taken from TestConstructingARayThroughCenterOfTheCanvas.
	c := NewCamera(201, 101, math.Pi/2)
	for i := 0; i < b.N; i++ {
		c.RayForPixel(100, 50)
	}
}

func BenchmarkCameraMethodRender(b *testing.B) {
	// Taken from TestRenderingWorldWithCamera.
	w := DefaultWorld()
	c := NewCamera(11, 11, math.Pi/2)
	from := NewPoint(0, 0, -5)
	to := NewPoint(0, 0, 0)
	up := NewVector(0, 1, 0)
	c.SetTransform(NewViewTransform(from, to, up))
	for i := 0; i < b.N; i++ {
		c.Render(w, 1, false)
	}
}
