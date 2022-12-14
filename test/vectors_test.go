// The test package ensures that all vectors functions are working
// as intended.
package test

import (
	"testing"

	vectors "github.com/jly02/linalg/pkg/vec"
)

const EPSILON float64 = 0.00000001

// Compares equality between two floats based on arbitrary epsilon value.
func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}

	return false
}

// Compare equality between two float arrays based on arbitrary epsilon value.
func floatSliceEquals(a, b []float64) bool {
	for i := range a {
		if !floatEquals(a[i], b[i]) {
			return false
		}
	}

	return true
}

func TestDotWithNil(t *testing.T) {
	var first []int = nil
	var second []int = nil

	want := 0.
	dot, err := vectors.Dot(first, second)
	if err != nil || !floatEquals(dot, want) {
		t.Fatalf(`Dot([], []) = %f, %v, want 0`, dot, err)
	}
}

func TestDotWithEmpty(t *testing.T) {
	first := make([]int, 0)
	second := make([]int, 0)

	want := 0.
	dot, err := vectors.Dot(first, second)
	if err != nil || !floatEquals(dot, want) {
		t.Fatalf(`Dot([0], [0]) = %f, %v, want 0`, dot, err)
	}
}

func TestDotWithOne(t *testing.T) {
	first := []int{5}
	second := []int{10}

	want := 50.
	dot, err := vectors.Dot(first, second)
	if err != nil || !floatEquals(dot, want) {
		t.Fatalf(`Dot([1], [1]) = %f, %v, want %f`, dot, err, want)
	}
}

func TestDotWithFew(t *testing.T) {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}

	want := 32.
	dot, err := vectors.Dot(first, second)
	if err != nil || !floatEquals(dot, want) {
		t.Fatalf(`Dot([3], [3]) = %f, %v, want %f`, dot, err, want)
	}
}

func TestDotWithMany(t *testing.T) {
	first := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	second := []int{11, 21, 31, 41, 51, 61, 71, 81, 91, 101, 111, 121, 131, 141, 151, 161, 171, 181, 191, 201}

	want := 28910.
	dot, err := vectors.Dot(first, second)
	if err != nil || !floatEquals(dot, want) {
		t.Fatalf(`Dot([20], [20]) = %f, %v, want %f`, dot, err, want)
	}
}

func TestDotWithNegative(t *testing.T) {
	first := []int{-1, -2, 3}
	second := []int{-3, 2, 4}

	want := 11.
	dot, err := vectors.Dot(first, second)
	if err != nil || !floatEquals(dot, want) {
		t.Fatalf(`Dot([3], [3]) = %f, %v, want %f`, dot, err, want)
	}
}

func TestDotWithDifferentLengths(t *testing.T) {
	first := make([]float64, 5)
	second := make([]int, 4)

	dot, err := vectors.Dot(first, second)
	if err == nil {
		t.Fatalf(`Dot([5], [4]) = %f, did not return error as expected`, dot)
	}
}

func TestDotWithOneNonInteger(t *testing.T) {
	first := []float64{.5, 2.5, 1.5}
	second := []int{1, 2, 4}

	want := 11.5
	dot, err := vectors.Dot(first, second)
	if err != nil || !floatEquals(dot, want) {
		t.Fatalf(`Dot([3], [3]) = %f, %v, want %f`, dot, err, want)
	}
}

func TestDotWithBothNonInteger(t *testing.T) {
	first := []float64{.5, 2.5, 1.5}
	second := []float64{1., 2., 4.}

	want := 11.5
	dot, err := vectors.Dot(first, second)
	if err != nil || !floatEquals(dot, want) {
		t.Fatalf(`Dot([3], [3]) = %f, %v, want %f`, dot, err, want)
	}
}

func TestIsOrthogonalOne(t *testing.T) {
	var first []int = nil
	var second []int = nil

	perp, err := vectors.IsOrthogonal(first, second)
	if err == nil {
		t.Fatalf(`IsOrthogonal(nil, nil) = %t, did not return error as expected`, perp)
	}
}

func TestIsOrthogonalFalse(t *testing.T) {
	first := []int{5}
	second := []int{10}

	want := false
	perp, err := vectors.IsOrthogonal(first, second)
	if err != nil || perp != want {
		t.Fatalf(`Dot(nil, nil) = %t, want %t, %s`, perp, want, err)
	}
}

func TestIsOrthogonalTrue(t *testing.T) {
	first := []int{15, -5, -5}
	second := []float64{5, 10, 5}

	want := true
	perp, err := vectors.IsOrthogonal(first, second)
	if err != nil || perp != want {
		t.Fatalf(`Dot(nil, nil) = %t, want %t, %s`, perp, want, err)
	}
}

func TestCrossWithNil(t *testing.T) {
	var first []int = nil
	var second []int = nil

	vec, err := vectors.Cross(first, second)
	if err == nil {
		t.Fatalf(`Cross(nil, nil) = %f didn't return error as expected`, vec)
	}
}

func TestCrossWithZero(t *testing.T) {
	first := []int{0, 0, 0}
	second := []int{0, 0, 0}

	want := []float64{0, 0, 0}
	vec, err := vectors.Cross(first, second)
	if err != nil || !floatSliceEquals(vec, want) {
		t.Fatalf(`Cross([3], [3]) = %f, want %f`, vec, want)
	}
}

func TestCrossWithSmallValues(t *testing.T) {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}

	want := []float64{-3, 6, -3}
	vec, err := vectors.Cross(first, second)
	if err != nil || !floatSliceEquals(vec, want) {
		t.Fatalf(`Cross([3], [3]) = %f, want %f`, vec, want)
	}
}

func TestCrossWithNegative(t *testing.T) {
	first := []int{-1, 2, 3}
	second := []int{4, 5, -6}

	want := []float64{-27, 6, -13}
	vec, err := vectors.Cross(first, second)
	if err != nil || !floatSliceEquals(vec, want) {
		t.Fatalf(`Cross([3], [3]) = %f, want %f`, vec, want)
	}
}

func TestCrossWithOneNonInteger(t *testing.T) {
	first := []float64{1.5, 2.5, 3.5}
	second := []int{4, 5, 6}

	want := []float64{-2.5, 5, -2.5}
	vec, err := vectors.Cross(first, second)
	if err != nil || !floatSliceEquals(vec, want) {
		t.Fatalf(`Cross([3], [3]) = %f, want %f`, vec, want)
	}
}

func TestCrossWithBothNonInteger(t *testing.T) {
	first := []float64{1.5, 2.5, 3.5}
	second := []float64{4.25, 5.65, 6.6}

	want := []float64{-3.275, 4.975, -2.15}
	vec, err := vectors.Cross(first, second)
	if err != nil || !floatSliceEquals(vec, want) {
		t.Fatalf(`Cross([3], [3]) = %f, want %f`, vec, want)
	}
}

func TestCrossWithWrongLength(t *testing.T) {
	first := []float64{1.5, 2.5, 3.5, 4.5}
	second := []float64{4.25, 5.65, 6.6}

	vec, err := vectors.Cross(first, second)
	if err == nil {
		t.Fatalf(`Cross([4], [3]) = %f didn't return error as expected`, vec)
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
	if err != nil || !floatSliceEquals(vec, want) {
		t.Fatalf(`Add([1], [1]) = {%f}, want {%f}`, vec[0], want)
	}
}

func TestAddWithFew(t *testing.T) {
	first := []int{1, 2, 3}
	second := []int{4, 5, 6}

	want := []float64{5, 7, 9}
	vec, err := vectors.Add(first, second)
	if err != nil || !floatSliceEquals(vec, want) {
		t.Fatalf(`Add([3], [3]) = {%f}, want {%f}`, vec, want)
	}
}

func TestAddWithMany(t *testing.T) {
	first := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	second := []int{20, 18, 16, 14, 12, 10, 8, 6, 4, 2}

	want := []float64{22, 22, 22, 22, 22, 22, 22, 22, 22, 22}
	vec, err := vectors.Add(first, second)
	if err != nil || !floatSliceEquals(vec, want) {
		t.Fatalf(`Add([10], [10]) = {%f}, want {%f}`, vec, want)
	}
}

func TestAddWithNegative(t *testing.T) {
	first := []int{-1, -2, -3}
	second := []int{3, 2, -1}

	want := []float64{2, 0, -4}
	vec, err := vectors.Add(first, second)
	if err != nil || !floatSliceEquals(vec, want) {
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

func TestAddWithOneNonInteger(t *testing.T) {
	first := []float64{.5, 2.5, 1.5}
	second := []int{1, 2, 4}

	want := []float64{1.5, 4.5, 5.5}
	vec, err := vectors.Add(first, second)
	if err != nil || !floatSliceEquals(vec, want) {
		t.Fatalf(`Add([3], [3]) = {%f}, want {%f}`, vec, want)
	}
}

func TestAddWithBothNonInteger(t *testing.T) {
	first := []float64{.5, 2.5, 1.5}
	second := []float64{1., 2., 4.}

	want := []float64{1.5, 4.5, 5.5}
	vec, err := vectors.Add(first, second)
	if err != nil || !floatSliceEquals(vec, want) {
		t.Fatalf(`Add([3], [3]) = {%f}, want {%f}`, vec, want)
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
	if err != nil || !floatSliceEquals(res, want) {
		t.Fatalf(`ScalarMult([3], 0) = {%f}, want {%f}`, res, want)
	}
}

func TestScalarMultWithOne(t *testing.T) {
	vec := []int{1}
	scalar := 3.

	want := []float64{3.}
	res, err := vectors.ScalarMult(vec, scalar)
	if err != nil || !floatSliceEquals(res, want) {
		t.Fatalf(`ScalarMult([1], 3) = {%f}, want {%f}`, res, want)
	}
}

func TestScalarMultWithMany(t *testing.T) {
	vec := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	scalar := 3.

	want := []float64{3., 6., 9., 12., 15., 18., 21., 24., 27., 30.}
	res, err := vectors.ScalarMult(vec, scalar)
	if err != nil || !floatSliceEquals(res, want) {
		t.Fatalf(`ScalarMult([10], 3) = {%f}, want {%f}`, res, want)
	}
}

func TestScalarMultWithNegative(t *testing.T) {
	vec := []int{1, 3, -4}
	scalar := -2.

	want := []float64{-2, -6, 8}
	res, err := vectors.ScalarMult(vec, scalar)
	if err != nil || !floatSliceEquals(res, want) {
		t.Fatalf(`ScalarMult([3], -2) = {%f}, want {%f}`, res, want)
	}
}

func TestScalarMultWithNonIntegerSlice(t *testing.T) {
	vec := []int{1, 2, 3, 4}
	scalar := .5

	want := []float64{.5, 1., 1.5, 2.}
	res, err := vectors.ScalarMult(vec, scalar)
	if err != nil || !floatSliceEquals(res, want) {
		t.Fatalf(`ScalarMult([4], .5) = {%f}, want {%f}`, res, want)
	}
}

func TestScalarMultWithNonIntegerScalar(t *testing.T) {
	vec := []float64{.5, 2., 3., 4.}
	scalar := 2

	want := []float64{1, 4, 6, 8}
	res, err := vectors.ScalarMult(vec, scalar)
	if err != nil || !floatSliceEquals(res, want) {
		t.Fatalf(`ScalarMult([4], 2) = {%f}, want {%f}`, res, want)
	}
}
