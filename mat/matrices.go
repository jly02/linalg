// The matrices package provides utility methods for generating matrices and
// performing standard linear-algebraic operations on them. All numeric results
// are returned as 64-bit floats, or  matrices of them, and all successful
// function calls will return err == nil.
package linalg

import "errors"

// Produces an identity matrix of a given size.
func Eye(size int) ([][]float64, error) {
	if size < 1 {
		return nil, errors.New("matrix cannot have size < 1")
	}

	mat := make([][]float64, size)
	for i := range mat {
		mat[i] = make([]float64, size)
		mat[i][i] = 1
	}

	return mat, nil
}
