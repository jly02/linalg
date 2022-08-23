// The vectors package provides utility methods
// for performing standard linear-algebraic
// operations on vectors.
package vectors

import "errors"

// Computes the dot product of two vectors.
//
// A successful Dot() call will return err == nil.
func Dot[T int | float64 | float32](first []T, second []T) (float64, error) {
	// Check if lengths are equal
	if len(first) != len(second) {
		return 0, errors.New("length of each vector must be equal")
	}

	var dot float64 = 0
	for i := 0; i < len(first); i++ {
		dot += float64(first[i]) * float64(second[i])
	}

	return dot, nil
}

// Adds two vectors, returning the resulting vector.
//
// A successful Add() call will return err == nil.
func Add[T int | float64 | float32](first []T, second []T) ([]float64, error) {
	// Check if lengths are equal
	if len(first) != len(second) {
		return nil, errors.New("length of each vector must be equal")
	}

	res := make([]float64, len(first))
	for i := 0; i < len(first); i++ {
		res[i] = float64(first[i]) + float64(second[i])
	}

	return res, nil
}

// Multiplies a given vector by a scalar value.
//
// A successful ScalarMult() call will return err == nil.
func ScalarMult[T int | float64 | float32](vec []T, scalar float64) ([]float64, error) {
	if len(vec) < 1 {
		return nil, errors.New("vector cannot be empty")
	}

	res := make([]float64, len(vec))
	if scalar == 0 {
		return res, nil
	}

	for i := 0; i < len(vec); i++ {
		res[i] = float64(vec[i]) * scalar
	}

	return res, nil
}
