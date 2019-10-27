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
	LocalPattern PatternInterface
	Transform    Matrix
}

func NewPattern(pi PatternInterface) Pattern {
	return Pattern{LocalPattern: pi, Transform: IdentityMatrix()}
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

	inversePatternTransform := p.Transform.Inverse()
	patternPoint := inversePatternTransform.MultiplyByTuple(objectPoint)

	return p.LocalPattern.LocalPatternAt(patternPoint)
}
