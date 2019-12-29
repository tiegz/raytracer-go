package raytracer

import (
	"fmt"
	"math"
)

type TextureMapPattern struct {
	Pattern Pattern
	UVMap   func(Tuple) (float64, float64)
}

func NewTextureMapPattern(p Pattern, f func(Tuple) (float64, float64)) Pattern {
	return NewPattern(TextureMapPattern{p, f})
}

func (p TextureMapPattern) String() string {
	return fmt.Sprintf("TextureMapPattern( %v )", p.Pattern) // TODO: add UVMap in here somehow
}

// Converts a point to spherical coordinates. (u = horizontal, v = vertical)
func SphericalMap(p Tuple) (float64, float64) {
	// compute the azimuthal angle (-π < theta <= π)
	// angle increases clockwise as viewed from above,
	// which is opposite of what we want, but we'll fix it later.
	theta := math.Atan2(p.X, p.Z)

	// vec is the vector pointing from the sphere's origin (the world origin)
	// to the point, which will also happen to be exactly equal to the sphere's
	// radius.
	vec := NewVector(p.X, p.Y, p.Z)
	radius := vec.Magnitude()

	// compute the polar angle (0 <= phi <= π)
	phi := math.Acos(p.Y / radius)

	//  -0.5 < raw_u <= 0.5
	rawU := theta / (2 * math.Pi)

	// 0 <= u < 1
	// here's also where we fix the direction of u. Subtract it from 1,
	// so that it increases counterclockwise as viewed from above.
	u := 1 - (rawU + 0.5)

	// we want v to be 0 at the south pole of the sphere,
	// and 1 at the north pole, so we have to "flip it over"
	// by subtracting it from 1.
	v := 1 - (phi / math.Pi)

	return u, v
}

/////////////////////////
// PatternInterface methods
/////////////////////////

// function pattern_at(texture_map, point)
//   let (u, v) ← texture_map.uv_map(point)
//   return uv_pattern_at(texture_map.uv_pattern, u, v)
// end function

func (p TextureMapPattern) LocalPatternAt(point Tuple) Color {
	u, v := p.UVMap(point)
	return p.Pattern.UVPatternAt(u, v)
}

func (p TextureMapPattern) LocalUVPatternAt(u, v float64) Color {
	return Colors["Black"]
}

func (cp TextureMapPattern) localIsEqualTo(cp2 PatternInterface) bool {
	// cp2Pattern := cp2.(*CheckerPattern)
	// if !cp.A.IsEqualTo(cp2Pattern.A) || !cp.B.IsEqualTo(cp2Pattern.B) {
	// 	return false
	// }
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (cp TextureMapPattern) localType() string {
	return "TextureMapPattern"
}
