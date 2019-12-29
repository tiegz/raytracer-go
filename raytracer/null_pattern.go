package raytracer

import (
	"fmt"
)

// TODO replace NullPattern pattern by just making Material.Pattern a pointer and checking against nil instead
type NullPattern struct {
}

func NewNullPattern() Pattern {
	return NewPattern(NullPattern{})
}

func (p NullPattern) String() string {
	return fmt.Sprintf("NullPattern( )")
}

/////////////////////////
// PatternInterface methods
/////////////////////////

func (np NullPattern) LocalPatternAt(point Tuple) Color {
	return Colors["White"]
}

func (np NullPattern) LocalUVPatternAt(u, v float64) Color {
	return Colors["Black"]
}

func (np NullPattern) localIsEqualTo(tp2 PatternInterface) bool {
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (np NullPattern) localType() string {
	return "NullPattern"
}
