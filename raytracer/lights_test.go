package raytracer

import (
	"fmt"
	"testing"
)

func TestPointLightHasPositionAndIntensity(t *testing.T) {
	i := Colors["White"]
	p := NewPoint(0, 0, 0)
	pl := NewPointLight(p, i)
	assertEqualTuple(t, p, pl.Corner)
	assertEqualColor(t, i, pl.GetIntensity())
}

func TestPointLightsEvaluateTheLightIntensityAtAGivenPoint(t *testing.T) {
	w := DefaultWorld()
	l := w.Lights[0]
	testCases := []struct {
		point  Tuple
		result float64
	}{
		{NewPoint(0, 1.0001, 0), 1.0},
		{NewPoint(-1.0001, 0, 0), 1.0},
		{NewPoint(0, 0, -1.0001), 1.0},
		{NewPoint(0, 0, 1.0001), 0.0},
		{NewPoint(1.0001, 0, 0), 0.0},
		{NewPoint(0, -1.0001, 0), 0.0},
		{NewPoint(0, 0, 0), 0.0},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("testCases[%d]", idx), func(t *testing.T) {
			assertEqualFloat64(t, tc.result, l.IntensityAt(tc.point, w))
		})
	}
}

func TestCreatingAnAreaLight(t *testing.T) {
	corner := NewPoint(0, 0, 0)
	v1 := NewVector(2, 0, 0)
	v2 := NewVector(0, 0, 1)
	light := NewAreaLight(corner, v1, 4, v2, 2, NewColor(1, 1, 1))

	assertEqualTuple(t, corner, light.Corner)
	assertEqualTuple(t, NewVector(0.5, 0, 0), light.UVec)
	assertEqualFloat64(t, 4, light.USteps)
	assertEqualTuple(t, NewVector(0, 0, 0.5), light.VVec)
	assertEqualFloat64(t, 2.0, light.VSteps)
	assertEqualFloat64(t, 8.0, light.Samples)
}

func TestFindingASinglePointOnAnAreaLight(t *testing.T) {
	light := NewAreaLight(NewPoint(0, 0, 0), NewVector(2, 0, 0), 4, NewVector(0, 0, 1), 2, NewColor(1, 1, 1))
	testCases := []struct {
		u      float64
		v      float64
		result Tuple
	}{
		{0, 0, NewPoint(0.25, 0, 0.25)},
		{1, 0, NewPoint(0.75, 0, 0.25)},
		{0, 1, NewPoint(0.25, 0, 0.75)},
		{2, 0, NewPoint(1.25, 0, 0.25)},
		{3, 1, NewPoint(1.75, 0, 0.75)},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("testCases[%d]", idx), func(t *testing.T) {
			assertEqualTuple(t, tc.result, light.PointOnLight(tc.u, tc.v))
		})
	}
}

func TestTheAreaLightIntensityFunction(t *testing.T) {
	w := DefaultWorld()
	light := NewAreaLight(NewPoint(-0.5, -0.5, -5), NewVector(1, 0, 0), 2, NewVector(0, 1, 0), 2, NewColor(1, 1, 1))
	testCases := []struct {
		Point  Tuple
		Result float64
	}{
		{NewPoint(0, 0, 2), 0.0},
		{NewPoint(1, -1, 2), 0.25},
		{NewPoint(1.5, 0, 2), 0.5},
		{NewPoint(1.25, 1.25, 3), 0.75},
		{NewPoint(0, 0, -2), 1.0},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("testCases[%d]", idx), func(t *testing.T) {
			assertEqualFloat64(t, tc.Result, light.IntensityAt(tc.Point, w))
		})
	}
}

func TestFindingAsinglePointOnAJitteredAreaLight(t *testing.T) {
	seq := NewSequence(0.3, 0.7)
	light := NewAreaLight(NewPoint(0, 0, 0), NewVector(2, 0, 0), 4, NewVector(0, 0, 1), 2, NewColor(1, 1, 1))
	light.Jitter = &seq
	testCases := []struct {
		u      float64
		v      float64
		Result Tuple
	}{
		{0, 0, NewPoint(0.15, 0, 0.35)},
		{1, 0, NewPoint(0.65, 0, 0.35)},
		{0, 1, NewPoint(0.15, 0, 0.85)},
		{2, 0, NewPoint(1.15, 0, 0.35)},
		{3, 1, NewPoint(1.65, 0, 0.85)},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("testCases[%d]", idx), func(t *testing.T) {
			pt := light.PointOnLight(tc.u, tc.v)
			assertEqualTuple(t, tc.Result, pt)
		})
	}
}

func TestTheAreaLightWithJitteredSamples(t *testing.T) {
	w := DefaultWorld()
	testCases := []struct {
		Point  Tuple
		Result float64
	}{
		{NewPoint(0, 0, 2), 0.0},
		{NewPoint(1, -1, 2), 0.5},
		{NewPoint(1.5, 0, 2), 0.75},
		{NewPoint(1.25, 1.25, 3), 0.75},
		{NewPoint(0, 0, -2), 1.0},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("testCases[%d]", idx), func(t *testing.T) {
			seq := NewSequence(0.7, 0.3, 0.9, 0.1, 0.5)
			light := NewAreaLight(NewPoint(-0.5, -0.5, -5), NewVector(1, 0, 0), 2, NewVector(0, 1, 0), 2, NewColor(1, 1, 1))
			light.Jitter = &seq
			intensity := light.IntensityAt(tc.Point, w)
			assertEqualFloat64(t, tc.Result, intensity)
		})
	}
}

/////////////
// Benchmarks
/////////////

func BenchmarkPointLightMethodIsEqualTo(b *testing.B) {
	light := NewPointLight(NewPoint(-0.5, -0.5, -5), NewColor(1, 1, 1))
	for i := 0; i < b.N; i++ {
		light.IsEqualTo(light)
	}
}

func BenchmarkAreaLightMethodIsEqualTo(b *testing.B) {
	light := NewAreaLight(NewPoint(-0.5, -0.5, -5), NewVector(1, 0, 0), 2, NewVector(0, 1, 0), 2, NewColor(1, 1, 1))
	for i := 0; i < b.N; i++ {
		light.IsEqualTo(light)
	}
}

func BenchmarkPointLightMethodIntensityAt(b *testing.B) {
	w := DefaultWorld()
	light := w.Lights[0]
	for i := 0; i < b.N; i++ {
		light.IntensityAt(NewPoint(0, 1.0001, 0), w)
	}
}

func BenchmarkAreaLightMethodIntensityAt(b *testing.B) {
	w := DefaultWorld()
	light := NewAreaLight(NewPoint(-0.5, -0.5, -5), NewVector(1, 0, 0), 2, NewVector(0, 1, 0), 2, NewColor(1, 1, 1))
	for i := 0; i < b.N; i++ {
		light.IntensityAt(NewPoint(0, 0, -2), w)
	}
}
