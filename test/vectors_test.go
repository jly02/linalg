package test

import (
	"reflect"
	"testing"

	vectors "github.com/jly02/linalg/pkg"
)

func TestDotWithNil(t *testing.T) {
	var first []int = nil
	var second []int = nil

	want := 0.
	dot, err := vectors.Dot(first, second)
	if err != nil || dot != want {
		t.Fatalf(`Dot([], []) = %f, %v, want 0`, dot, err)
	}
}

func TestDotWithEmpty(t *testing.T) {
	first := make([]int, 0)
	second := make([]int, 0)

	want := 0.
	dot, err := vectors.Dot(first, second)
	if err != nil || dot != want {
		t.Fatalf(`Dot([0], [0]) = %f, %v, want 0`, dot, err)
	}
}

func TestDotWithOne(t *testing.T) {
	first := []int{5}
	second := []int{10}

	want := 50.
	dot, err := vectors.Dot(first, second)
	if err != nil || dot != want {
		t.Fatalf(`Dot([1], [1]) = %f, %v, want 0`, dot, err)
	}
}

func TestDotWithFew(t *testing.T) {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}

	want := 32.
	dot, err := vectors.Dot(first, second)
	if err != nil || dot != want {
		t.Fatalf(`Dot([3], [3]) = %f, %v, want 0`, dot, err)
	}
}

func TestDotWithMany(t *testing.T) {
	first := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	second := []int{11, 21, 31, 41, 51, 61, 71, 81, 91, 101, 111, 121, 131, 141, 151, 161, 171, 181, 191, 201}

	want := 28910.
	dot, err := vectors.Dot(first, second)
	if err != nil || dot != want {
		t.Fatalf(`Dot([20], [20]) = %f, %v, want 0`, dot, err)
	}
}

func TestDotWithNegative(t *testing.T) {
	first := []int{-1, -2, 3}
	second := []int{-3, 2, 4}

	want := 11.
	dot, err := vectors.Dot(first, second)
	if err != nil || dot != want {
		t.Fatalf(`Dot([3], [3]) = %f, %v, want 0`, dot, err)
	}
}

func TestDotWithDifferentLengths(t *testing.T) {
	first := make([]int, 5)
	second := make([]int, 4)

	dot, err := vectors.Dot(first, second)
	if err == nil {
		t.Fatalf(`Dot([5], [4]) = %f, did not return error as expected`, dot)
	}
}

func TestAddWithNil(t *testing.T) {
	var first []int = nil
	var second []int = nil

	vec, err := vectors.Add(first, second)
	if err != nil {
		t.Fatalf(`Add(nil, nil) = %f returned error`, vec)
	}
}

func TestAddWithOne(t *testing.T) {
	first := []int{1}
	second := []int{2}

	want := []float64{3}
	vec, err := vectors.Add(first, second)
	if err != nil || !reflect.DeepEqual(want, vec) {
		t.Fatalf(`Add([1], [1]) = {%f}, want {%f}`, vec[0], want)
	}
}

func TestAddWithFew(t *testing.T) {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}

	want := []float64{5, 7, 9}
	vec, err := vectors.Add(first, second)
	if err != nil || !reflect.DeepEqual(vec, want) {
		t.Fatalf(`Add([3], [3]) = {%f}, want {%f}`, vec, want)
	}
}

func TestAddWithMany(t *testing.T) {
	first := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	second := []int{20, 18, 16, 14, 12, 10, 8, 6, 4, 2}

	want := []float64{22, 22, 22, 22, 22, 22, 22, 22, 22, 22}
	vec, err := vectors.Add(first, second)
	if err != nil || !reflect.DeepEqual(vec, want) {
		t.Fatalf(`Add([10], [10]) = {%f}, want {%f}`, vec, want)
	}
}

func TestAddWithNegative(t *testing.T) {
	first := []int{-1, -2, -3}
	second := []int{3, 2, -1}

	want := []float64{2, 0, -4}
	vec, err := vectors.Add(first, second)
	if err != nil || !reflect.DeepEqual(vec, want) {
		t.Fatalf(`Add([3], [3]) = {%f}, want {%f}`, vec, want)
	}
}

func TestAddWithDifferentLengths(t *testing.T) {
	first := make([]int, 5)
	second := make([]int, 4)

	vec, err := vectors.Add(first, second)
	if err == nil {
		t.Fatalf(`Add([5], [4]) = %f, did not return error as expected`, vec)
	}
}

func TestScalarMultWithNil(t *testing.T) {
	var vec []int = nil
	scalar := 0.

	res, err := vectors.ScalarMult(vec, scalar)
	if err == nil {
		t.Fatalf(`ScalarMult(nil, 0) = %f, did not return error as expected`, res)
	}
}

func TestScalarMultWithZeroScalar(t *testing.T) {
	vec := []int{1, 2, 3}
	scalar := 0.

	want := []float64{0., 0., 0.}
	res, err := vectors.ScalarMult(vec, scalar)
	if err != nil || !reflect.DeepEqual(res, want) {
		t.Fatalf(`ScalarMult([3], 0) = {%f}, want {%f}`, res, want)
	}
}

func TestScalarMultWithOne(t *testing.T) {
	vec := []int{1}
	scalar := 3.

	want := []float64{3.}
	res, err := vectors.ScalarMult(vec, scalar)
	if err != nil || !reflect.DeepEqual(res, want) {
		t.Fatalf(`ScalarMult([1], 3) = {%f}, want {%f}`, res, want)
	}
}

func TestScalarMultWithMany(t *testing.T) {
	vec := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	scalar := 3.

	want := []float64{3., 6., 9., 12., 15., 18., 21., 24., 27., 30.}
	res, err := vectors.ScalarMult(vec, scalar)
	if err != nil || !reflect.DeepEqual(res, want) {
		t.Fatalf(`ScalarMult([10], 3) = {%f}, want {%f}`, res, want)
	}
}

func TestScalarMultWithNegative(t *testing.T) {
	vec := []int{1, 3, -4}
	scalar := -2.

	want := []float64{-2, -6, 8}
	res, err := vectors.ScalarMult(vec, scalar)
	if err != nil || !reflect.DeepEqual(res, want) {
		t.Fatalf(`ScalarMult([3], -2) = {%f}, want {%f}`, res, want)
	}
}

func TestScalarMultWithNonInteger(t *testing.T) {
	vec := []int{1, 2, 3, 4}
	scalar := .5

	want := []float64{.5, 1., 1.5, 2.}
	res, err := vectors.ScalarMult(vec, scalar)
	if err != nil || !reflect.DeepEqual(res, want) {
		t.Fatalf(`ScalarMult([3], -2) = {%f}, want {%f}`, res, want)
	}
}
