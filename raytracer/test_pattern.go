package raytracer

import (
	"fmt"
)

// This is just a dummy "child" of Pattern, for testing purposes.
type TestPattern struct {
}

func NewTestPattern() *Pattern {
	return NewPattern(TestPattern{})
}

func (p TestPattern) String() string {
	return fmt.Sprintf("TestPattern( )")
}

/////////////////////////
// PatternInterface methods
/////////////////////////

func (tp TestPattern) LocalPatternAt(point Tuple) Color {
	return NewColor(point.X, point.Y, point.Z)
}

func (tp TestPattern) LocalUVPatternAt(u, v float64) Color {
	return Colors["Black"]
}

func (tp TestPattern) localIsEqualTo(tp2 PatternInterface) bool {
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (tp TestPattern) localType() string {
	return "TestPattern"
}
