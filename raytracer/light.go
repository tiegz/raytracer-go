package raytracer

import "fmt"

type AreaLight struct {
	Corner    Tuple   // corner: position of one corner of the light source
	UVec      Tuple   // direction+length of the u edge
	USteps    float64 // how many points are sampled along u edge. More steps = less banding, but with jittering it becomes noisier.
	VVec      Tuple   // direction+length of the v edge
	VSteps    float64 // how many points are sampled along v edge. More steps = less banding, but with jittering it becomes noisier.
	Intensity Color   // intensity of the light
	Samples   float64
	Jitter    *Sequence
}

// Returns a flat, rectangular light source -- composed of cells -- that casts a soft shadow.
//    |----------------------------|
// ^  |     |     |     |     |    |
// |  |----------------------------|
// v  |     |     |  ^  |     |    |
// v  |----------------------------|
// e  |     |  <  |cells|  >  |    |
// c  |____________________________|
// corner       uvec ->
//
// TODO: investigate adaptive subdivision, to algorithically determine the resolution (i.e. steps) to use:
//   https://pdfs.semanticscholar.org/9792/e5563ac82ad33ffd6c9c0772682e96d6ba72.pdf
//   http://www.cse.chalmers.se/~uffe/xjobb/SoftShadows2.pdf
func NewAreaLight(corner Tuple, full_uvec Tuple, usteps float64, full_vvec Tuple, vsteps float64, intensity Color) *AreaLight {
	jitter := NewSequence(0.5)
	al := &AreaLight{
		Corner:    corner,
		UVec:      full_uvec.Divide(usteps),
		USteps:    usteps,
		VVec:      full_vvec.Divide(vsteps),
		VSteps:    vsteps,
		Samples:   usteps * vsteps,
		Intensity: intensity,
		Jitter:    &jitter,
	}
	return al
}

// Returns a simple point light, composed of 1 cell.
func NewPointLight(position Tuple, intensity Color) *AreaLight {
	pl := NewAreaLight(position, NewVector(1, 0, 0), 1, NewVector(0, 1, 0), 1, intensity)
	jitter := NewSequence(0.0) // PointLights are single points of light and don't need any jiter.
	pl.Jitter = &jitter
	return pl
}

func (al *AreaLight) String() string {
	return fmt.Sprintf(
		"AreaLight(\nCorner: %v\nUVec: %v\nUSteps: %v\nVVec: %v\nVSteps: %v\nIntensity: %v\nSamples: %v\nJitter: %v\n)",
		al.Corner,
		al.UVec,
		al.USteps,
		al.VVec,
		al.VSteps,
		al.Intensity,
		al.Samples,
		al.Jitter,
	)
}

func (al *AreaLight) GetIntensity() Color {
	return al.Intensity
}

func (al *AreaLight) IsEqualTo(al2 *AreaLight) bool {
	if !al.Corner.IsEqualTo(al2.Corner) {
		return false
	} else if !al.UVec.IsEqualTo(al2.UVec) {
		return false
	} else if al.USteps != al2.USteps {
		return false
	} else if !al.VVec.IsEqualTo(al2.VVec) {
		return false
	} else if al.VSteps != al2.VSteps {
		return false
	} else if al.Samples != al2.Samples {
		return false
	} else if !al.Jitter.IsEqualTo(*al2.Jitter) {
		return false
	}
	return true
}

func (al AreaLight) IntensityAt(p Tuple, w *World) float64 {
	total := 0.0
	for v := 0.0; v < al.VSteps; v++ {
		for u := 0.0; u < al.USteps; u++ {
			lightPosition := al.PointOnLight(u, v)
			if !w.IsShadowed(lightPosition, p) {
				total += 1.0
			}
		}
	}
	return total / al.Samples
}

// Returns the real point on the area light based on the u/v coordinates.
// This places the position of the point randomly based on the Jitter sequence
// to avoid the banding produced by a uniform position.
func (al *AreaLight) PointOnLight(u, v float64) Tuple {
	return al.Corner.
		Add(al.UVec.Multiply(u + al.Jitter.Next())).
		Add(al.VVec.Multiply(v + al.Jitter.Next()))
}
