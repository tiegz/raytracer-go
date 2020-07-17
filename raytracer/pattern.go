package raytracer

import (
	"fmt"
	"math"
)

type PatternInterface interface {
	LocalPatternAt(Tuple) Color
	LocalUVPatternAt(float64, float64) Color // TODO: only used for uv patterns -- should this just be a different interface?
	localIsEqualTo(PatternInterface) bool
	localType() string
}

// Pattern is a general pattern (Transform), with the specific type of pattern stored as a PatternInterface in LocalPattern.
type Pattern struct {
	LocalPattern     PatternInterface
	Transform        Matrix // WARNING: don't set Transform directly, use SetTransform()
	InverseTransform Matrix
}

func NewPattern(pi PatternInterface) *Pattern {
	p := Pattern{LocalPattern: pi}
	p.SetTransform(IdentityMatrix())
	return &p
}

func (p *Pattern) SetTransform(m Matrix) {
	p.Transform = m
	p.InverseTransform = m.Inverse()
}

func (p *Pattern) String() string {
	return fmt.Sprintf("Pattern(\n  LocalPattern: %v\n  Transform: %v\n)", p.LocalPattern, p.Transform)
}

func (p *Pattern) IsEqualTo(p2 *Pattern) bool {
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

func (p *Pattern) PatternAtShape(s Shape, worldPoint Tuple) Color {
	objectPoint := s.WorldToObject(worldPoint)
	patternPoint := p.InverseTransform.MultiplyByTuple(objectPoint)

	return p.LocalPattern.LocalPatternAt(patternPoint)
}

func (p *Pattern) UVPatternAt(u, v float64) Color {
	return p.LocalPattern.LocalUVPatternAt(u, v)
}

// TODO: is there a better value to return than string? Maybe enum/iota?
// Returns which face a given point on a unit cube is on.
func FaceFromPoint(p Tuple) string {
	absX := math.Abs(p.X)
	absY := math.Abs(p.Y)
	absZ := math.Abs(p.Z)
	coord := maxFloat64(absX, absY, absZ)

	switch coord {
	case p.X:
		return "right"
	case -p.X:
		return "left"
	case p.Y:
		return "up"
	case -p.Y:
		return "down"
	case p.Z:
		return "front"
	default:
		return "back"
	}
}

// Maps a point on the front face of a cube to its uv values.
func CubeUVFront(p Tuple) (float64, float64) {
	u := math.Mod((p.X+1), 2.0) / 2.0
	v := math.Mod((p.Y+1), 2.0) / 2.0

	return u, v
}

// Maps a point on the back face of a cube to its uv values.
func CubeUVBack(p Tuple) (float64, float64) {
	u := math.Mod((1-p.X), 2.0) / 2.0
	v := math.Mod((p.Y+1), 2.0) / 2.0

	return u, v
}

// Maps a point on the left face of a cube to its uv values.
func CubeUVLeft(p Tuple) (float64, float64) {
	u := math.Mod((p.Z+1), 2.0) / 2.0
	v := math.Mod((p.Y+1), 2.0) / 2.0

	return u, v
}

// Maps a point on the right face of a cube to its uv values.
func CubeUVRight(p Tuple) (float64, float64) {
	u := math.Mod((1-p.Z), 2.0) / 2.0
	v := math.Mod((p.Y+1), 2.0) / 2.0

	return u, v
}

// Maps a point on the upper face of a cube to its uv values.
func CubeUVUpper(p Tuple) (float64, float64) {
	u := math.Mod((p.X+1), 2.0) / 2.0
	v := math.Mod((1-p.Z), 2.0) / 2.0

	return u, v
}

func CubeUVLower(p Tuple) (float64, float64) {
	u := math.Mod((p.X+1), 2.0) / 2.0
	v := math.Mod((p.Z+1), 2.0) / 2.0

	return u, v
}
