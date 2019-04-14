package raytracer

import (
	"fmt"
	"math"
)

type Pattern struct {
	A         Color
	B         Color
	Transform Matrix
}

func NullPattern() Pattern {
	return Pattern{}
}

func NewStripePattern(a, b Color) Pattern {
	return Pattern{a, b, IdentityMatrix()}
}

func (p Pattern) String() string {
	return fmt.Sprintf("Pattern( )")
}

func (p *Pattern) IsEqualTo(p2 Pattern) bool {
	if !p.A.IsEqualTo(p2.A) || !p.B.IsEqualTo(p2.B) || !p.Transform.IsEqualTo(p2.Transform) {
		return false
	}
	return true
}

func (p *Pattern) StripeAt(point Tuple) Color {
	if math.Mod(math.Floor(point.X), 2) == 0 {
		return p.A
	} else {
		return p.B
	}
}

func (p *Pattern) StripeAtObject(obj Shape, point Tuple) Color {
	inverseObjTransform := obj.Transform.Inverse()
	objectPoint := inverseObjTransform.MultiplyByTuple(point)

	inversePatternTransform := p.Transform.Inverse()
	patternPoint := inversePatternTransform.MultiplyByTuple(objectPoint)

	return p.StripeAt(patternPoint)
}
