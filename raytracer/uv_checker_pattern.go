package raytracer

import (
	"fmt"
	"math"
)

type UVCheckerPattern struct {
	Width  float64
	Height float64
	A      Color
	B      Color
}

type TextureMap struct {
	pattern  Pattern
	function func(Tuple) (float64, float64)
}

func NewTextureMap(p Pattern, f func(Tuple) (float64, float64)) TextureMap {
	return TextureMap{p, f}
}

func NewUVCheckerPattern(w, h float64, a, b Color) Pattern {
	return NewPattern(UVCheckerPattern{w, h, a, b})
}

func (p UVCheckerPattern) String() string {
	return fmt.Sprintf("UVCheckerPattern( Width: %v Height: %v %v %v )", p.Width, p.Height, p.A, p.B)
}

/////////////////////////
// PatternInterface methods
/////////////////////////

func (p UVCheckerPattern) LocalPatternAt(point Tuple) Color {
	return Colors["Black"]
}

func (p UVCheckerPattern) LocalUVPatternAt(u, v float64) Color {
	u2 := math.Floor(u * p.Width)
	v2 := math.Floor(v * p.Height)

	if int(u2+v2)%2 == 0 {
		return p.A
	} else {
		return p.B
	}
}

func (p UVCheckerPattern) localIsEqualTo(p2 PatternInterface) bool {
	p2Pattern := p2.(*UVCheckerPattern)
	if !p.A.IsEqualTo(p2Pattern.A) || !p.B.IsEqualTo(p2Pattern.B) {
		return false
	}
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (p UVCheckerPattern) localType() string {
	return "UVCheckerPattern"
}
