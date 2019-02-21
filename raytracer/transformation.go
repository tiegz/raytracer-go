package raytracer

import "math"

func NewTranslation(x, y, z float64) Matrix {
	translation := IdentityMatrix()
	translation.Set(0, 3, x)
	translation.Set(1, 3, y)
	translation.Set(2, 3, z)
	return translation
}

func NewScale(x, y, z float64) Matrix {
	scale := NewMatrix(4, 4, []float64{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	})
	return scale
}

// aka "skew"
func NewShear(xToY, xToZ, yToX, yToZ, zToX, zToY float64) Matrix {
	scale := NewMatrix(4, 4, []float64{
		1, xToY, xToZ, 0,
		yToX, 1, yToZ, 0,
		zToX, zToY, 1, 0,
		0, 0, 0, 1,
	})
	return scale
}

func NewRotateX(radians float64) Matrix {
	scale := NewMatrix(4, 4, []float64{
		1, 0, 0, 0,
		0, math.Cos(radians), -math.Sin(radians), 0,
		0, math.Sin(radians), math.Cos(radians), 0,
		0, 0, 0, 1,
	})
	return scale
}

func NewRotateY(radians float64) Matrix {
	scale := NewMatrix(4, 4, []float64{
		math.Cos(radians), 0, math.Sin(radians), 0,
		0, 1, 0, 0,
		-math.Sin(radians), 0, math.Cos(radians), 0,
		0, 0, 0, 1,
	})
	return scale
}

func NewRotateZ(radians float64) Matrix {
	scale := NewMatrix(4, 4, []float64{
		math.Cos(radians), -math.Sin(radians), 0, 0,
		math.Sin(radians), math.Cos(radians), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	})
	return scale
}
