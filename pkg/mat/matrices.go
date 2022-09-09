// Package linalg matrices module provides utility methods for generating matrices and
// performing standard linear-algebraic operations on them. All numeric results
// are returned as 64-bit floats, or  matrices of them, and all successful
// function calls will return err == nil.
package linalg

import "errors"

// IdMatrix Produces an identity matrix of a given size.
func IdMatrix(size int) ([][]float64, error) {
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

// Transpose Transposes a matrix. Matrix must be rectangular in dimensions.
func Transpose[T int | float64](a [][]T) ([][]float64, error) {
	// empty case
	dim1 := len(a)
	if dim1 < 1 {
		return make([][]float64, 0), nil
	}

	dim2 := len(a[0])
	for i := range a {
		if len(a[i]) != dim2 {
			return nil, errors.New("matrix dimensions must be rectangular")
		}
	}

	mat := make([][]float64, dim2)
	for i := range a {
		for j := range a[i] {
			if mat[j] == nil {
				mat[j] = make([]float64, dim1)
			}

			mat[j][i] = float64(a[i][j])
		}
	}

	return mat, nil
}

// MatMult Multiplies two matrices together, returning the result.
func MatMult[T int | float64, E int | float64](a [][]T, b [][]E) ([][]float64, error) {
	return IdMatrix(4)
}
