package raytracer

import (
	"fmt"
)

type CubeMapPattern struct {
	left  Pattern
	front Pattern
	right Pattern
	back  Pattern
	upper Pattern
	lower Pattern
}

func NewCubeMapPattern(l, f, r, b, u, d Pattern) Pattern {
	return NewPattern(CubeMapPattern{l, f, r, b, u, d})
}

func (p CubeMapPattern) String() string {
	return fmt.Sprintf("CubeMapPattern( %v %v %v %v %v %v )", p.left, p.front, p.right, p.back, p.upper, p.lower)
}

/////////////////////////
// PatternInterface methods
/////////////////////////

func (p CubeMapPattern) LocalPatternAt(point Tuple) Color {
	var u, v float64
	face := FaceFromPoint(point)

	switch face {
	case "left":
		u, v = CubeUVLeft(point)
		return p.left.UVPatternAt(u, v)
	case "right":
		u, v = CubeUVRight(point)
		return p.right.UVPatternAt(u, v)
	case "front":
		u, v = CubeUVFront(point)
		return p.front.UVPatternAt(u, v)
	case "back":
		u, v = CubeUVBack(point)
		return p.back.UVPatternAt(u, v)
	case "up":
		u, v = CubeUVUpper(point)
		return p.upper.UVPatternAt(u, v)
	default: // down
		u, v = CubeUVLower(point)
		return p.lower.UVPatternAt(u, v)
	}
}

func (p CubeMapPattern) LocalUVPatternAt(u, v float64) Color {
	return Colors["Black"]
}

func (p CubeMapPattern) localIsEqualTo(p2 PatternInterface) bool {
	p2Pattern := p2.(*CubeMapPattern)
	if !p.left.IsEqualTo(p2Pattern.left) {
		return false
	} else if !p.right.IsEqualTo(p2Pattern.right) {
		return false
	} else if !p.upper.IsEqualTo(p2Pattern.upper) {
		return false
	} else if !p.lower.IsEqualTo(p2Pattern.lower) {
		return false
	} else if !p.front.IsEqualTo(p2Pattern.front) {
		return false
	} else if !p.back.IsEqualTo(p2Pattern.back) {
		return false
	}
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (p CubeMapPattern) localType() string {
	return "CubeMapPattern"
}
