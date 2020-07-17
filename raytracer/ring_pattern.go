package raytracer

import (
	"fmt"
	"math"
)

type RingPattern struct {
	A Color
	B Color
}

func NewRingPattern(a, b Color) *Pattern {
	return NewPattern(RingPattern{a, b})
}

func (p RingPattern) String() string {
	return fmt.Sprintf("RingPattern( A: %v B: %v )", p.A, p.B)
}

/////////////////////////
// PatternInterface methods
/////////////////////////

func (p RingPattern) LocalPatternAt(point Tuple) Color {
	xSquared := math.Pow(point.X, 2)
	zSquared := math.Pow(point.Z, 2)
	if int(math.Floor(math.Sqrt(xSquared+zSquared)))%2 == 0 {
		return p.A
	} else {
		return p.B
	}
}

func (p RingPattern) LocalUVPatternAt(u, v float64) Color {
	return Colors["Black"]
}

func (rp RingPattern) localIsEqualTo(rp2 PatternInterface) bool {
	rp2Pattern := rp2.(*RingPattern)
	if !rp.A.IsEqualTo(rp2Pattern.A) || !rp.B.IsEqualTo(rp2Pattern.B) {
		return false
	}
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (rp RingPattern) localType() string {
	return "RingPattern"
}
