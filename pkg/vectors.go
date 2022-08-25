// The vectors package provides utility methods for performing standard
// linear-algebraic operations on vectors. All numeric results are returned
// as 64-bit floats, or vectors of them.
package vectors

import (
	"errors"
	"math"
)

const EPSILON float64 = 0.00000001

// Computes the dot product of two vectors.
//
// A successful Dot() call will return err == nil.
func Dot[T int | float64, E int | float64](a []T, b []E) (float64, error) {
	// Check if lengths are equal
	if len(a) != len(b) {
		return 0, errors.New("length of each vector must be equal")
	}

	var dot float64 = 0
	for i := 0; i < len(a); i++ {
		dot += float64(a[i]) * float64(b[i])
	}

	return dot, nil
}

// Determines whether two vectors are orthogonal, I.E. their dot product is zero,
// with an error within EPSILON.
//
// A successful isOrthogonal() call will return err == nil.
func IsOrthogonal[T int | float64, E int | float64](a []T, b []E) (bool, error) {
	if a == nil || b == nil {
		return false, errors.New("vectors cannot be nil")
	}

	dot, err := Dot(a, b)
	return math.Abs(dot-0) < EPSILON, err
}

// Computes the cross product of two vectors. A cross product is exclusive to 3-dimensional
// vectors.
//
// A successful Cross() call will return err == nil.
func Cross[T int | float64, E int | float64](a []T, b []E) ([]float64, error) {
	// Both vectors must be 3D
	if len(a) != 3 || len(b) != 3 {
		return nil, errors.New("length of both vectors must be 3")
	}

	res := make([]float64, 3)
	res[0] = float64(a[1])*float64(b[2]) - float64(a[2])*float64(b[1]) // (a2 * b3) - (a3 * b2)
	res[1] = float64(a[2])*float64(b[0]) - float64(a[0])*float64(b[2]) // (a3 * b1) - (a1 * b3)
	res[2] = float64(a[0])*float64(b[1]) - float64(a[1])*float64(b[0]) // (a1 * b2) - (a2 * b1)

	return res, nil
}

// Adds two vectors, returning the resulting vector.
//
// A successful Add() call will return err == nil.
func Add[T int | float64, E int | float64](a []T, b []E) ([]float64, error) {
	// Check if lengths are equal
	if len(a) != len(b) {
		return nil, errors.New("length of each vector must be equal")
	}

	res := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = float64(a[i]) + float64(b[i])
	}

	return res, nil
}

// Multiplies a given vector by a scalar value.
//
// A successful ScalarMult() call will return err == nil.
func ScalarMult[T int | float64, E int | float64](vec []T, scalar E) ([]float64, error) {
	if len(vec) < 1 {
		return nil, errors.New("vector cannot be empty")
	}

	res := make([]float64, len(vec))
	if scalar == 0 {
		return res, nil
	}

	for i := 0; i < len(vec); i++ {
		res[i] = float64(vec[i]) * float64(scalar)
	}

	return res, nil
}
