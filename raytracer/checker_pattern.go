package raytracer

import (
	"fmt"
	"math"
)

type CheckerPattern struct {
	A Color
	B Color
}

func NewCheckerPattern(a, b Color) Pattern {
	return NewPattern(CheckerPattern{a, b})
}

func (p CheckerPattern) String() string {
	return fmt.Sprintf("CheckerPattern( %v %v )", p.A, p.B)
}

/////////////////////////
// PatternInterface methods
/////////////////////////

func (p CheckerPattern) LocalPatternAt(point Tuple) Color {
	if int(math.Abs(point.X)+math.Abs(point.Y)+math.Abs(point.Z))%2 == 0 {
		return p.A
	} else {
		return p.B
	}
}

func (cp CheckerPattern) localIsEqualTo(cp2 PatternInterface) bool {
	cp2Pattern := cp2.(*CheckerPattern)
	if !cp.A.IsEqualTo(cp2Pattern.A) || !cp.B.IsEqualTo(cp2Pattern.B) {
		return false
	}
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (cp CheckerPattern) localType() string {
	return "CheckerPattern"
}