package examples

import (
	"fmt"
	"io/ioutil"
	"math"

	. "github.com/tiegz/raytracer-go/pkg/raytracer"
)

func RunDrawAnalogClockExample() {
	canvas_width := 200
	canvas_halfwidth := float64(canvas_width / 2)
	clock_radius_scale := 0.4 // as a ratio to canvas width
	clock_radius := float64(canvas_width) * clock_radius_scale
	canvas := NewCanvas(canvas_width+1, canvas_width+1)
	translate_center := NewTranslation(canvas_halfwidth, canvas_halfwidth, 0)

	// Draw the center point.
	c := NewPoint(0, 1, 0)
	center := translate_center.MultiplyByTuple(c)
	canvas.WritePixel(int(math.Round(center.X)), int(math.Round(center.Y)), NewColor(1, 0, 0))

	// Drawgst each hour's point.
	for i := 0.0; i < 12.0; i += 1.0 {
		transformation := NewTranslation(canvas_halfwidth, canvas_halfwidth, 0)           // center the point
		transformation = transformation.Multiply(NewRotateZ((math.Pi * float64(i) / 6)))  // rotate to the hour's position
		transformation = transformation.Multiply(NewScale(clock_radius, clock_radius, 0)) // scale the point relative to canvas size

		point := transformation.MultiplyByTuple(c)
		x := int(math.Round(point.X))
		y := int(math.Round(point.Y))
		canvas.WritePixel(x, y, NewColor(12-i/12, i/12+0.1, 0))
	}

	fmt.Println("Generating PPM...")
	ppm := canvas.ToPpm()
	filename := "tmp/analog_clock.ppm"
	ppmBytes := []byte(ppm)
	fmt.Printf("Saving scene to %s...\n", filename)
	if err := ioutil.WriteFile(filename, ppmBytes, 0644); err != nil {
		panic(err)
	}
}
