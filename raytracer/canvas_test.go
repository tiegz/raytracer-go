package raytracer

import (
	"testing"
)

func TestNewCanvas(t *testing.T) {
  c1 := NewCanvas(10, 20)

  assertEqualFloat64(t, 10, c1.Width)
  assertEqualFloat64(t, 20, c1.Height)
}

