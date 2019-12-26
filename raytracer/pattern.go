package raytracer

import (
	"fmt"
)

type PatternInterface interface {
	LocalPatternAt(Tuple) Color
	localIsEqualTo(PatternInterface) bool
	localType() string
}

// Pattern is a general pattern (Transform), with the specific type of pattern stored as a PatternInterface in LocalPattern.
type Pattern struct {
	LocalPattern     PatternInterface
	Transform        Matrix // WARNING: don't set Transform directly, use SetTransform()
	InverseTransform Matrix
}

func NewPattern(pi PatternInterface) Pattern {
	p := Pattern{LocalPattern: pi}
	p.SetTransform(IdentityMatrix())
	return p
}

func (p *Pattern) SetTransform(m Matrix) {
	p.Transform = m
	p.InverseTransform = m.Inverse()
}

func (p Pattern) String() string {
	return fmt.Sprintf("Pattern( %v )", p.LocalPattern)
}

func (p Pattern) IsEqualTo(p2 Pattern) bool {
	pt1 := p.LocalPattern.localType()
	pt2 := p2.LocalPattern.localType()

	if pt1 != pt2 {
		return false
	} else if !p.Transform.IsEqualTo(p2.Transform) {
		return false
	} else {
		return p.LocalPattern.localIsEqualTo(p2.LocalPattern)
	}
}

func (p Pattern) PatternAtShape(s Shape, worldPoint Tuple) Color {
	objectPoint := s.WorldToObject(worldPoint)
	patternPoint := p.InverseTransform.MultiplyByTuple(objectPoint)

	return p.LocalPattern.LocalPatternAt(patternPoint)
}
