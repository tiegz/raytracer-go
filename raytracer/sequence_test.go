package raytracer

import (
	"testing"
)

func TestANumberGeneratorReturnsACyclicSequenceOfNumbers(t *testing.T) {
	gen := NewSequence(0.1, 0.5, 1.0)
	assertEqualFloat64(t, 0.1, gen.Next())
	assertEqualFloat64(t, 0.5, gen.Next())
	assertEqualFloat64(t, 1.0, gen.Next())
	assertEqualFloat64(t, 0.1, gen.Next())
}
