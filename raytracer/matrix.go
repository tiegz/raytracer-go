package raytracer

import "fmt"

type Matrix struct {
	Rows int
	Cols int
	Data []float64
}

func NewMatrix(rows, cols int, data []float64) Matrix {
	return Matrix{rows, cols, data}
}

func (m *Matrix) At(row, col int) float64 {
	return m.Data[row*m.Cols+col]
}

func (m *Matrix) Set(row, col int, val float64) {
	m.Data[row*m.Cols+col] = val
}

func (m *Matrix) IsEqualTo(m2 Matrix) bool {
	if m.Cols != m2.Cols || m.Rows != m2.Rows {
		return false
	}

	for col := 0; col < m.Cols; col = col + 1 {
		for row := 0; row < m.Rows; row = row + 1 {
			if m.At(row, col) != m2.At(row, col) {
				return false
			}
		}
	}

	return true
}

func IdentityMatrix() Matrix {
	// Only implementing a 4x4 identity matrix for now.
	return NewMatrix(4, 4, []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	})
}

func (m *Matrix) Transpose() Matrix {
	m2 := NewMatrix(m.Rows, m.Cols, make([]float64, m.Rows*m.Cols, m.Rows*m.Cols))

	for r := 0; r < m.Rows; r += 1 {
		for c := 0; c < m.Cols; c += 1 {
			m2.Set(r, c, m.At(c, r))
		}
	}

	return m2
}

func (m *Matrix) Multiply(m2 Matrix) Matrix {
	m3 := NewMatrix(m.Rows, m.Cols, make([]float64, m.Rows*m.Cols, m.Rows*m.Cols))

	for row := 0; row < m.Rows; row += 1 {
		for col := 0; col < m.Cols; col += 1 {

			// Note: this would only work with square matrices. (pg 29)
			sum := 0.0
			for i := 0; i < m.Rows; i += 1 {
				sum += m.At(row, i) * m2.At(i, col)
			}

			m3.Set(row, col, sum)
		}
	}

	return m3
}

func (m *Matrix) MutiplyByTuple(t Tuple) Tuple {
	// Similar to mulitplying matrices, except there's only one col in a Tuple.
	x := m.At(0, 0)*t.X +
		m.At(0, 1)*t.Y +
		m.At(0, 2)*t.Z +
		m.At(0, 3)*t.W

	y := m.At(1, 0)*t.X +
		m.At(1, 1)*t.Y +
		m.At(1, 2)*t.Z +
		m.At(1, 3)*t.W

	z := m.At(2, 0)*t.X +
		m.At(2, 1)*t.Y +
		m.At(2, 2)*t.Z +
		m.At(2, 3)*t.W

	w := m.At(3, 0)*t.X +
		m.At(3, 1)*t.Y +
		m.At(3, 2)*t.Z +
		m.At(3, 3)*t.W

	return Tuple{x, y, z, w}
}

// .. If the determinant is zero, then the corresponding system of equations has no solution...
func (m *Matrix) Determinant() float64 {
	if m.Rows == 2 {
		return (m.At(0, 0) * m.At(1, 1)) - (m.At(1, 0) * m.At(0, 1))
	} else {
		// The magical thing is that it doesn’t matter which row or column you choose. It just works.
		row, sum := 0, 0.0
		for col := 0; col < m.Cols; col += 1 {
			sum += m.At(row, col) * m.Cofactor(row, col)
		}
		return sum
	}
}

// Returns the determinant of a submatrix.
func (m *Matrix) Minor(rowToRemove, colToRemove int) float64 {
	sub := m.Submatrix(rowToRemove, colToRemove)
	return sub.Determinant()
}

// ...Cofactors are minors that have (possibly) had their sign changed...
func (m *Matrix) Cofactor(rowToRemove, colToRemove int) float64 {
	minor := m.Minor(rowToRemove, colToRemove)
	if (rowToRemove+colToRemove)%2 == 0 {
		return minor
	} else {
		return -minor
	}
}

func (m *Matrix) Submatrix(rowToRemove, colToRemove int) Matrix {
	r, c := m.Rows-1, m.Cols-1
	m2 := NewMatrix(r, c, make([]float64, r*c))

	fmt.Printf("Submatrix \n")
	for rowOrig, rowNew := 0, 0; rowOrig < m.Rows; rowOrig += 1 {
		if rowOrig != rowToRemove {
			for colOrig, colNew := 0, 0; colOrig < m.Cols; colOrig += 1 {
				if colOrig != colToRemove {
					m2.Set(rowNew, colNew, m.At(rowOrig, colOrig))
					colNew += 1
				}
			}
			rowNew += 1
		}
	}

	return m2
}
