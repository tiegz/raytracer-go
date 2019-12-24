package raytracer

import (
	"testing"
)

func TestNewColor(t *testing.T) {
	t1 := NewColor(-0.5, 0.4, 1.7)

	assertEqualFloat64(t, -0.5, t1.Red)
	assertEqualFloat64(t, 0.4, t1.Green)
	assertEqualFloat64(t, 1.7, t1.Blue)
}

func TestAddingTwoColors(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	expected := NewColor(1.6, 0.7, 1.0)
	actual := c1.Add(c2)

	assertEqualColor(t, expected, actual)
}

func TestSubtractingTwoColors(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	expected := NewColor(0.2, 0.5, 0.5)
	actual := c1.Subtract(c2)

	assertEqualColor(t, expected, actual)
}

func TestMultiplyingColorByScalar(t *testing.T) {
	c1 := NewColor(0.2, 0.3, 0.4)
	expected := NewColor(0.4, 0.6, 0.8)
	actual := c1.Multiply(2)

	assertEqualColor(t, expected, actual)
}

func TestMultiplyingColors(t *testing.T) {
	c1 := NewColor(1, 0.2, 0.4)
	c2 := NewColor(0.9, 1, 0.1)
	expected := NewColor(0.9, 0.2, 0.04)
	actual := c1.MultiplyColor(c2)

	assertEqualColor(t, expected, actual)
}

/////////////
// Benchmarks
/////////////

func BenchmarkColorMethodIsEqualTo(b *testing.B) {
	c1 := NewColor(0.8, 0.1, 0.3)
	for i := 0; i < b.N; i++ {
		c1.IsEqualTo(c1)
	}
}
