package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	"github.com/tiegz/raytracer-go/raytracer"
)

func RunDrawAnalogClockExample() {
	canvas_width := 200
	canvas_halfwidth := float64(canvas_width / 2)
	point_scale := 0.4
	canvas_scale := float64(canvas_width) * point_scale
	canvas := raytracer.NewCanvas(canvas_width+1, canvas_width+1)
	translate_center := raytracer.NewTranslation(canvas_halfwidth, canvas_halfwidth, 0)

	// Draw the center point.
	c := raytracer.NewPoint(0, 1, 0)
	center := translate_center.MultiplyByTuple(c)
	canvas.WritePixel(int(math.Round(center.X)), int(math.Round(center.Y)), raytracer.NewColor(1, 0, 0))

	// Drawgst each hour's point.
	for i := 0.0; i < 12.0; i += 1.0 {
		transformation := raytracer.NewTranslation(canvas_halfwidth, canvas_halfwidth, 0)           // center the point
		transformation = transformation.Multiply(raytracer.NewRotateZ((math.Pi * float64(i) / 6)))  // rotate to the hour's position
		transformation = transformation.Multiply(raytracer.NewScale(canvas_scale, canvas_scale, 0)) // scale the point relative to canvas size

		point := transformation.MultiplyByTuple(c)
		x := int(math.Round(point.X))
		y := int(math.Round(point.Y))
		canvas.WritePixel(x, y, raytracer.NewColor(12-i/12, i/12+0.1, 0))
	}

	fmt.Println("Generating PPM...")
	ppm := canvas.ToPpm()
	filename := "tmp/analog_clock.ppm"
	ppmBytes := []byte(ppm)
	fmt.Printf("Saving analog clock to %s...\n", filename)
	if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
		panic(err)
	}
}
