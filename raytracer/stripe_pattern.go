package raytracer

import (
	"fmt"
	"math"
)

// This is just a dummy "child" of Pattern, for testing purposes.
type StripePattern struct {
	A Color
	B Color
}

func NewStripePattern(a, b Color) Pattern {
	return NewPattern(StripePattern{a, b})
}

func (s StripePattern) String() string {
	return fmt.Sprintf("StripePattern( %v %v )", s.A, s.B)
}

/////////////////////////
// PatternInterface methods
/////////////////////////

func (s StripePattern) LocalPatternAt(point Tuple) Color {
	if math.Mod(math.Floor(point.X), 2) == 0 {
		return s.A
	} else {
		return s.B
	}
}

func (sp StripePattern) localIsEqualTo(sp2 PatternInterface) bool {
	sp2StripePattern := sp2.(*StripePattern)
	if !sp.A.IsEqualTo(sp2StripePattern.A) || !sp.B.IsEqualTo(sp2StripePattern.B) {
		return false
	}
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (sp StripePattern) localType() string {
	return "StripePattern"
}
