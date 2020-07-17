package raytracer

import (
	"fmt"
	"math"
)

type GradientPattern struct {
	A Color
	B Color
}

func NewGradientPattern(a, b Color) *Pattern {
	return NewPattern(GradientPattern{a, b})
}

func (p GradientPattern) String() string {
	return fmt.Sprintf("GradientPattern( A: %v B: %v )", p.A, p.B)
}

/////////////////////////
// PatternInterface methods
/////////////////////////

// ... [a blending function] is a function that takes two values and interpolates the values between them ...
func (p GradientPattern) LocalPatternAt(point Tuple) Color {
	// color(p, ca, cb) = ca + (cb − ca) ∗ (px − floor(px))
	distance := p.B.Subtract(p.A)
	fraction := point.X - math.Floor(point.X)
	return p.A.Add(distance.Multiply(fraction))
}

func (p GradientPattern) LocalUVPatternAt(u, v float64) Color {
	return Colors["Black"]
}

func (gp GradientPattern) localIsEqualTo(gp2 PatternInterface) bool {
	gp2Pattern := gp2.(*GradientPattern)
	if !gp.A.IsEqualTo(gp2Pattern.A) || !gp.B.IsEqualTo(gp2Pattern.B) {
		return false
	}
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (gp GradientPattern) localType() string {
	return "GradientPattern"
}
