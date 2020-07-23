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

func NewUVAlignCheckPattern(main, ul, ur, bl, br Color) *Pattern {
	return NewPattern(UVAlignCheckPattern{main, ul, ur, bl, br})
}

func (acp UVAlignCheckPattern) String() string {
	return fmt.Sprintf(
		"AlignCheckPattern( %s %s %s %s %s )",
		acp.main,
		acp.ul,
		acp.ur,
		acp.bl,
		acp.br,
	)
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

func (acp UVAlignCheckPattern) localIsEqualTo(acp2 PatternInterface) bool {
	acp2Pattern := acp2.(*UVAlignCheckPattern)
	if !acp.main.IsEqualTo(acp2Pattern.main) {
		return false
	} else if !acp.ul.IsEqualTo(acp2Pattern.ul) {
		return false
	} else if !acp.ur.IsEqualTo(acp2Pattern.ur) {
		return false
	} else if !acp.bl.IsEqualTo(acp2Pattern.bl) {
		return false
	} else if !acp.br.IsEqualTo(acp2Pattern.br) {
		return false
	}
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (acp UVAlignCheckPattern) localType() string {
	return "UVAlignCheckPattern"
}
