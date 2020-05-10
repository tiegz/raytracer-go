package raytracer

import (
	"fmt"
)

// This is a pattern for testing purposes.
type UVAlignCheckPattern struct {
	main Color
	ul   Color
	ur   Color
	bl   Color
	br   Color
}

func NewUVAlignCheckPattern(main, ul, ur, bl, br Color) Pattern {
	return NewPattern(UVAlignCheckPattern{main, ul, ur, bl, br})
}

func (acp UVAlignCheckPattern) String() string {
	return fmt.Sprintf("AlignCheckPattern( )")
}

/////////////////////////
// PatternInterface methods
/////////////////////////

func (acp UVAlignCheckPattern) LocalPatternAt(point Tuple) Color {
	return NewColor(point.X, point.Y, point.Z)
}

func (acp UVAlignCheckPattern) LocalUVPatternAt(u, v float64) Color {
	if v > 0.8 {
		if u < 0.2 {
			return acp.ul
		}
		if u > 0.8 {
			return acp.ur
		}
	} else if v < 0.2 {
		if u < 0.2 {
			return acp.bl
		}
		if u > 0.8 {
			return acp.br
		}
	}

	return acp.main
}

func (acp UVAlignCheckPattern) localIsEqualTo(tp2 PatternInterface) bool {
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (acp UVAlignCheckPattern) localType() string {
	return "UVAlignCheckPattern"
}
