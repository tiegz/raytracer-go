package raytracer

import (
	"fmt"
	"math"
)

// This is a pattern for testing purposes.
type UVImagePattern struct {
	Canvas Canvas
}

func NewUVImagePattern(c Canvas) Pattern {
	return NewPattern(UVImagePattern{c})
}

func (ip UVImagePattern) String() string {
	return fmt.Sprintf("UVImagePattern( )")
}

/////////////////////////
// PatternInterface methods
/////////////////////////

func (ip UVImagePattern) LocalPatternAt(point Tuple) Color {
	return NewColor(point.X, point.Y, point.Z)
}

func (ip UVImagePattern) LocalUVPatternAt(u, v float64) Color {
	v = 1 - v

	x := u * float64(ip.Canvas.Width-1)
	y := v * float64(ip.Canvas.Height-1)

	return ip.Canvas.PixelAt(int(math.Round(x)), int(math.Round(y)))
}

func (ip UVImagePattern) localIsEqualTo(tp2 PatternInterface) bool {
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (ip UVImagePattern) localType() string {
	return "UVImagePattern"
}
