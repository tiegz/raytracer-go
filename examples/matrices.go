package examples

import (
	"fmt"

	"github.com/tiegz/raytracer-go/raytracer"
)

func RunMatrixExample() {
	m := raytracer.NewMatrix(3, 3, []float64{
		1, 2, 3,
		5, 6, 7,
		8, 9, 0,
	})
	fmt.Printf("Given the matrix:\n%v\n", m)
	im := m.Inverse()
	fmt.Printf("The inverted matrix is:\n%v\n", im)
	tm := m.Transpose()
	fmt.Printf("The transposed matrix is:\n%v\n", tm)
	mim := m.Multiply(im)
	fmt.Printf("The matrix multipled by its inverse is:\n%v\n", mim)

	fmt.Println()
	identityMatrix := raytracer.IdentityMatrix()
	fmt.Printf("Given the identiy matrix:\n%v\n", identityMatrix)
	fmt.Printf("The inverted matrix is:\n%v\n", identityMatrix.Inverse())
	fmt.Printf("The transposed matrix is:\n%v\n", identityMatrix.Transpose())
}
