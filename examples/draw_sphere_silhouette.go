package examples

import (
	"fmt"
	"math"

	. "github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawSphereSilhouette(scaleX, scaleY, rotateZ, skew bool) {
	canvasSize := 250
	wallZ := 10.0
	wallSize := 7.0
	halfWallSize := float64(wallSize) / 2.0
	pixelSize := wallSize / float64(canvasSize)

	rayOrigin := NewPoint(0, 0, -5)
	canvas := NewCanvas(int(canvasSize), int(canvasSize))
	sphere := NewSphere()
	transform := IdentityMatrix()
	if scaleX {
		scale := NewScale(1, 0.5, 1)
		transform = scale.Multiply(transform)
	}
	if scaleY {
		scale := NewScale(0.5, 1, 1)
		transform = scale.Multiply(transform)
	}
	if rotateZ {
		rotate := NewRotateZ(math.Pi / 4)
		transform = rotate.Multiply(transform)
	}
	if skew {
		shear := NewShear(1, 0, 0, 0, 0, 0)
		transform = shear.Multiply(transform)
	}
	sphere.SetTransform(transform)

	// for each row of pixels in the canvas
	for y := 0; y < canvasSize; y++ {
		// compute the world y coordinate (top = +half, bottom = -half)
		worldY := halfWallSize - (pixelSize * float64(y))
		// for each col of pixels in the canvas
		for x := 0; x < canvasSize; x++ {
			// compute the world x coordinate (left = -half, right = half)
			worldX := -halfWallSize + (pixelSize * float64(x))

			// describe the point on the wall that the ray will target
			targetPoint := NewPoint(worldX, worldY, wallZ)
			rayDirection := targetPoint.Subtract(rayOrigin)
			r := NewRay(rayOrigin, rayDirection)
			intersections := sphere.Intersect(r)

			if hit := intersections.Hit(false); !hit.IsNull() {
				canvas.WritePixel(int(x), int(y), Colors["Purple"])
			}
		}
	}

	if err := canvas.SaveJPEG("tmp/world.jpg"); err != nil {
		fmt.Printf("Something went wrong! %s\n", err)
	} else {
		fmt.Println("Saved to tmp/world.jpg")
	}
}
