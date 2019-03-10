package raytracer

import "math"

// NewViewTransform returns a transformation matrix that orients the world relative
// to your eye.
//
// The from argument is the position of the eye. (Point)
// The to argument is the position the eye's looking at. (Point)
// The up argument specifies the direction of up. (Vector)
func NewViewTransform(from, to, up Tuple) Matrix {
	fwd := to.Subtract(from)
	fwd = fwd.Normalized()

	left := fwd.Cross(up.Normalized())

	trueUp := left.Cross(fwd)

	orientation := NewMatrix(4, 4, []float64{
		left.X, left.Y, left.Z, 0,
		trueUp.X, trueUp.Y, trueUp.Z, 0,
		-fwd.X, -fwd.Y, -fwd.Z, 0,
		0, 0, 0, 1,
	})

	return orientation.Multiply(NewTranslation(-from.X, -from.Y, -from.Z))
}

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
