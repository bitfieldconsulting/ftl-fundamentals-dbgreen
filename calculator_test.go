package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add (%f,%f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 1, b: 2, want: -1},
		{a: 100, b: 99, want: 1},
		{a: 3, b: 3, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f,%f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}
func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: -1, b: 5, want: -5},
		{a: 9, b: 3, want: 27},
		{a: 5, b: 5, want: 25},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f,%f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}
func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b        float64
		want        float64
		errExpected bool
	}
	testCases := []testCase{
		{a: 6, b: 0, want: 0, errExpected: true},
		{a: 6, b: 3, want: 2, errExpected: false},
		{a: 10, b: 4, want: 2.5, errExpected: false},
	}

	for _, tc := range testCases {
		got, e := calculator.Divide(tc.a, tc.b)
		if tc.errExpected == true && e == nil {
			t.Errorf("Divide (%f,%f): %s", tc.a, tc.b, e)
		} else if tc.errExpected == false && e != nil {
			t.Errorf("Divide(%f,%f): Unexpected Error", tc.a, tc.b)
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("Divide (%f,%f): wanted %f, got %f", tc.a, tc.b, tc.want, got)
		}

	}
}
