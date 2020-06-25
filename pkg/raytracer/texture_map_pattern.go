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
	return fmt.Sprintf("TextureMapPattern(\n  %v\n  %T\n)", p.Pattern, p.UVMap)
}

// Converts a point to spherical coordinates. (u = horizontal, v = vertical)
func SphericalMap(p Tuple) (float64, float64) {
	// the azimuthal angle (-π < theta <= π)
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

// Converts a point to planar coordinates. (u = horizontal, v = vertical)
func PlanarMap(p Tuple) (float64, float64) {
	u := altMod(p.X, 1)
	v := altMod(p.Z, 1)

	return u, v
}

// Converts a point to planar coordinates. (u = horizontal, v = vertical)
// TODO solve problem of top/bottom being rendered based on y component
func CylindricalMap(p Tuple) (float64, float64) {
	// ... compute the azimuthal angle, same as with spherical_map() ...
	theta := math.Atan2(p.X, p.Z) // should be p.Y?
	rawU := theta / (2 * math.Pi)
	u := 1 - (rawU + 0.5)

	// let v go from 0 to 1 between whole units of y
	v := altMod(p.Y, 1)

	return u, v
}

/////////////////////////
// PatternInterface methods
/////////////////////////

func (p TextureMapPattern) LocalPatternAt(point Tuple) Color {
	// TODO: is this right?
	u, v := p.UVMap(point)
	return p.Pattern.UVPatternAt(u, v)
}

func (p TextureMapPattern) LocalUVPatternAt(u, v float64) Color {
	return p.Pattern.UVPatternAt(u, v)
}

func (p TextureMapPattern) localIsEqualTo(p2 PatternInterface) bool {
	p2Pattern := p2.(*TextureMapPattern)
	if !p.Pattern.IsEqualTo(p2Pattern.Pattern) {
		return false
		// TODO: add IsEqualTo to UVMap
		// } else if p.UVMap.IsEqualTo(p2Pattern.UVMap) {
		// 	return false
	}
	return true
}

// Not returning reflect.TypeOf here because I suspect it
// does the same thing under the hood and stores a string?
func (p TextureMapPattern) localType() string {
	return "TextureMapPattern"
}
